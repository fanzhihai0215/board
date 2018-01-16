package auth

import (
	"git/inspursoft/board/src/apiserver/service"
	"git/inspursoft/board/src/common/model"
	"git/inspursoft/board/src/common/utils"
	"os"
	"testing"
	"fmt"
	
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
)

const (
	adminUserID     = 1
	initialPassword = "123456a?"
)

func updateAdminPassword() {
	salt := utils.GenerateRandomString()
	encryptedPassword := utils.Encrypt(initialPassword, salt)
	user := model.User{ID: adminUserID, Password: encryptedPassword, Salt: salt}
	isSuccess, err := service.UpdateUser(user, "password", "salt")
	if err != nil {
		logs.Error("Failed to update user password: %+v", err)
	}
	if isSuccess {
		logs.Info("Admin password has been updated successfully.")
	} else {
		logs.Info("Failed to update admin initial password.")
	}
}

func connectToDB() {
	hostIP:=os.Getenv("HOST_IP") 
	err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("root:root123@tcp(%s:3306)/board?charset=utf8", hostIP))
	if err != nil {
		logs.Error("Failed to connect to DB.")
	}
}

func TestMain(m *testing.M) {
	connectToDB()
	updateAdminPassword()
	os.Exit(m.Run())
}

func TestSignIn(t *testing.T) {
	assert := assert.New(t)
	currentAuth, err := GetAuth("db_auth")
	u, err := (*currentAuth).DoAuth("admin", "123456a?")
	assert.Nil(err, "Error occurred while calling SignIn method.")
	assert.NotNil(u, "User is nil.")
	assert.Equal("admin", u.Username, "Signed in failed.")
}

func TestSignInLdap(t *testing.T) {
	hostIP:=os.Getenv("HOST_IP")
	utils.Initialize()
	utils.SetConfig("LDAP_URL", fmt.Sprintf("ldap://%s", hostIP))
	utils.SetConfig("LDAP_SEARCH_DN", `cn=admin,dc=example,dc=org`)
	utils.SetConfig("LDAP_BASE_DN", "uid=test,dc=example,dc=org")
	utils.SetConfig("LDAP_FILTER", "")
	utils.SetConfig("LDAP_SEARCH_PWD", "admin")
	utils.SetConfig("LDAP_UID", "cn")
	utils.SetConfig("LDAP_SCOPE", "LDAP_SCOPE_SUBTREE")
	utils.SetConfig("LDAP_SCOPE", "")
	utils.SetConfig("LDAP_TIMEOUT", "5")
	assert := assert.New(t)
	currentAuth, err := GetAuth("ldap_auth")
	u, err := (*currentAuth).DoAuth(`test`, `123456`)
	assert.Nil(err, "Error occurred while calling SignIn method.")
	assert.NotNil(u, "User is nil.")
}
