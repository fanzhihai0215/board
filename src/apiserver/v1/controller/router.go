package controller

import (
	"git/inspursoft/board/src/apiserver/controllers/commons"

	"github.com/astaxie/beego"
)

func InitRouter() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/sign-in",
				&AuthController{},
				"post:SignInAction"),
			beego.NSRouter("/captcha",
				&commons.CaptchaController{}),
			beego.NSRouter("/cache-store",
				&commons.CacheStoreController{}),
			beego.NSRouter("/ext-auth",
				&AuthController{},
				"get:ExternalAuthAction"),
			beego.NSRouter("/sign-up",
				&AuthController{},
				"post:SignUpAction"),
			beego.NSRouter("/log-out",
				&AuthController{},
				"get:LogOutAction"),
			beego.NSRouter("/user-exists",
				&AuthController{},
				"get:UserExists"),
			beego.NSRouter("/users/current",
				&AuthController{},
				"get:CurrentUserAction"),
			beego.NSRouter("/systeminfo",
				&AuthController{},
				"get:GetSystemInfo"),
			beego.NSRouter("system/resources",
				&AuthController{},
				"get:GetSystemResources"),
			beego.NSRouter("/kubernetesinfo",
				&AuthController{},
				"get:GetKubernetesInfo"),
			beego.NSRouter("/reset-password",
				&AuthController{},
				"post:ResetPassword"),
			beego.NSRouter("/users",
				&UserController{},
				"get:GetUsersAction"),
			beego.NSRouter("/users/:id([0-9]+)/password",
				&UserController{},
				"put:ChangePasswordAction"),
			beego.NSRouter("/users/changeaccount",
				&UserController{},
				"put:ChangeUserAccount"),
			beego.NSRouter("/adduser",
				&SystemAdminController{},
				"post:AddUserAction"),
			beego.NSRouter("/users/:id([0-9]+)",
				&SystemAdminController{},
				"get:GetUserAction;put:UpdateUserAction;delete:DeleteUserAction"),
			beego.NSRouter("/users/:id([0-9]+)/systemadmin",
				&SystemAdminController{},
				"put:ToggleSystemAdminAction"),
			beego.NSRouter("/projects",
				&ProjectController{},
				"head:ProjectExists;get:GetProjectsAction;post:CreateProjectAction"),
			beego.NSRouter("/projects/:id([0-9]+)/publicity",
				&ProjectController{},
				"put:ToggleProjectPublicAction"),
			beego.NSRouter("/projects/:id([0-9]+)",
				&ProjectController{},
				"get:GetProjectAction;delete:DeleteProjectAction"),
			beego.NSRouter("/projects/:id([0-9]+)/members",
				&ProjectMemberController{},
				"get:GetProjectMembersAction;post:AddOrUpdateProjectMemberAction"),
			beego.NSRouter("/projects/:projectId([0-9]+)/members/:userId([0-9]+)",
				&ProjectMemberController{},
				"delete:DeleteProjectMemberAction"),
			beego.NSRouter("/images",
				&ImageController{},
				"get:GetImagesAction;delete:DeleteImageAction"),
			beego.NSRouter("/images/:imagename(.*)",
				&ImageController{},
				"get:GetImageDetailAction;delete:DeleteImageTagAction"),
			beego.NSRouter("/images/building",
				&ImageController{},
				"post:BuildImageAction"),
			beego.NSRouter("/images/imagepackage",
				&ImageController{},
				"post:UploadAndPushImagePackageAction"),
			beego.NSRouter("/images/dockerfilebuilding",
				&ImageController{},
				"post:DockerfileBuildImageAction"),
			beego.NSRouter("/images/dockerfile",
				&ImageController{},
				"get:GetImageDockerfileAction"),
			beego.NSRouter("/images/registry",
				&ImageController{},
				"get:GetImageRegistryAction"),
			beego.NSRouter("/images/preview",
				&ImageController{},
				"post:DockerfilePreviewAction"),
			beego.NSRouter("/images/configclean",
				&ImageController{},
				"delete:ConfigCleanAction"),
			beego.NSRouter("/images/checkused",
				&ImageController{},
				"get:GetImageUsedAction"),
			beego.NSRouter("/images/:imagename(.*)/existing",
				&ImageController{},
				"get:CheckImageTagExistingAction"),
			beego.NSRouter("/search",
				&SearchSourceController{}, "get:Search"),
			beego.NSRouter("/pvolumes",
				&PVolumeController{},
				"get:GetPVolumeListAction;post:AddPVolumeAction"),
			beego.NSRouter("/pvolumes/:id([0-9]+)",
				&PVolumeController{},
				"get:GetPVolumeAction;delete:RemovePVolumeAction"),
			beego.NSRouter("/pvolumes/existing",
				&PVolumeController{},
				"get:CheckPVolumeNameExistingAction"),
			beego.NSRouter("/pvclaims",
				&PVClaimController{},
				"get:GetPVClaimListAction;post:AddPVClaimAction"),
			beego.NSRouter("/pvclaims/:id([0-9]+)",
				&PVClaimController{},
				"get:GetPVClaimAction;delete:RemovePVClaimAction"),
			beego.NSRouter("/pvclaims/existing",
				&PVClaimController{},
				"get:GetPVCNameExisting"),
			beego.NSRouter("/configmaps",
				&ConfigMapController{},
				"get:GetConfigMapListAction;post:AddConfigMapAction"),
			beego.NSRouter("/configmaps/:configmapname([\\w]+)",
				&ConfigMapController{},
				"get:GetConfigMapAction;delete:RemoveConfigMapAction;put:UpdateConfigMapAction"),
			beego.NSRouter("/node",
				&NodeController{}, "get:GetNode"),
			beego.NSRouter("/nodes",
				&NodeController{}, "get:NodeList;post:AddNodeAction"),
			beego.NSRouter("/nodes/:nodename(.*)",
				&NodeController{}, "get:GetNodeStatusAction;delete:RemoveNodeAction"),
			beego.NSRouter("/nodes/:nodename(.*)/drain",
				&NodeController{}, "put:NodeDrainAction"),
			beego.NSRouter("/nodes/availableresources",
				&NodeController{}, "get:NodesAvailalbeResources"),
			beego.NSRouter("/node/toggle",
				&NodeController{}, "get:NodeToggle"),
			beego.NSRouter("/node/:id([0-9]+)/group",
				&NodeController{},
				"get:GetGroupsOfNodeAction;post:AddNodeToGroupAction;delete:RemoveNodeFromGroupAction"),
			beego.NSRouter("/edgenodes",
				&NodeController{}, "get:EdgeNodeList;post:AddEdgeNodeAction"),
			beego.NSRouter("/edgenodes/:nodename(.*)",
				&NodeController{}, "get:GetEdgeNodeAction;delete:RemoveEdgeNodeAction"),
			beego.NSRouter("/nodegroup",
				&NodeGroupController{},
				"get:GetNodeGroupsAction;post:AddNodeGroupAction;delete:DeleteNodeGroupAction"),
			beego.NSRouter("/nodegroup/:id([0-9]+)",
				&NodeGroupController{},
				"delete:DeleteNodeGroupAction"),
			beego.NSRouter("/nodegroup/existing",
				&NodeGroupController{},
				"get:CheckNodeGroupNameExistingAction"),
			beego.NSNamespace("/storage",
				beego.NSRouter("/setnfs", &StorageController{}, "post:Storage"),
			),
			beego.NSNamespace("/dashboard", beego.NSRouter("/service",
				&DashboardServiceController{},
				"post:GetServiceData"),
				beego.NSRouter("/node",
					&DashboardNodeController{}, "post:GetNodeData"),
				beego.NSRouter("/adminservercheck",
					&Dashboard{}, "get:AdminserverCheck"),
				beego.NSRouter("/checksysbyadminserver",
					&Dashboard{}, "get:CheckSysByAdminserver"),
				beego.NSRouter("/data",
					&Dashboard{}, "post:GetData"),
				beego.NSRouter("/time",
					&DashboardServiceController{}, "get:GetServerTime"),
			),
			beego.NSRouter("/services",
				&ServiceController{},
				"post:CreateServiceConfigAction;get:GetServiceListAction"),
			beego.NSRouter("/services/exists",
				&ServiceController{},
				"get:ServiceExists"),
			beego.NSRouter("/services/rollingupdate/session",
				&ServiceRollingUpdateController{},
				"get:GetServiceSessionFlagAction;patch:PatchServiceSessionAction"),
			beego.NSRouter("/services/rollingupdate/image",
				&ServiceRollingUpdateController{},
				"get:GetRollingUpdateServiceImageConfigAction;patch:PatchRollingUpdateServiceImageAction"),
			beego.NSRouter("/services/rollingupdate/nodegroup",
				&ServiceRollingUpdateController{},
				"get:GetRollingUpdateServiceNodeGroupConfigAction;patch:PatchRollingUpdateServiceNodeGroupAction"),
			beego.NSRouter("/services/:id([0-9]+)",
				&ServiceController{},
				"delete:DeleteServiceAction"),
			beego.NSRouter("/services/:id([0-9]+)/deployment",
				&ServiceController{},
				"delete:DeleteDeployAction"),
			beego.NSRouter("/services/deployment",
				&ServiceController{},
				"post:DeployServiceAction"),
			beego.NSRouter("/services/statefulsets",
				&ServiceController{},
				"post:DeployStatefulSetAction"),
			beego.NSRouter("/services/:id([0-9]+)/statefulsets",
				&ServiceController{},
				"delete:DeleteStatefulSetAction"),
			beego.NSRouter("/services/config",
				&ServiceConfigController{},
				"post:SetConfigServiceStepAction;get:GetConfigServiceStepAction;delete:DeleteServiceStepAction"),
			beego.NSRouter("/services/:id([0-9]+)/status",
				&ServiceController{},
				"get:GetServiceStatusAction"),
			beego.NSRouter("/services/selectservices",
				&ServiceController{},
				"get:GetSelectableServicesAction"),
			beego.NSRouter("/services/yaml/upload",
				&ServiceController{},
				"post:UploadYamlFileAction"),
			beego.NSRouter("/services/yaml/download",
				&ServiceController{},
				"get:DownloadDeploymentYamlFileAction"),
			beego.NSRouter("/services/nodeports",
				&ServiceController{},
				"get:GetServiceNodePorts"),
			beego.NSRouter("/services/import",
				&ServiceController{},
				"get:ImportServicesAction"),
			beego.NSRouter("/images/dockerfile/upload",
				&ImageController{},
				"post:UploadDockerfileFileAction"),
			beego.NSRouter("/images/dockerfile/download",
				&ImageController{},
				"get:DownloadDockerfileFileAction"),
			beego.NSRouter("/services/:id([0-9]+)/info",
				&ServiceController{},
				"get:GetServiceInfoAction"),
			beego.NSRouter("/services/:id([0-9]+)/logs/:podname(.+)",
				&ServiceController{},
				"get:GetServicePodLogsAction"),
			beego.NSRouter("/services/info",
				&ServiceController{},
				"post:StoreServiceRoute"),
			beego.NSRouter("/services/:id([0-9]+)/test",
				&ServiceController{},
				"post:DeployServiceTestAction"),
			beego.NSRouter("/services/:id([0-9]+)/toggle",
				&ServiceController{},
				"put:ToggleServiceAction"),
			beego.NSRouter("/services/:id([0-9]+)/scale",
				&ServiceController{},
				"put:ScaleServiceAction;get:GetScaleStatusAction"),
			beego.NSRouter("/services/:id([0-9]+)/publicity",
				&ServiceController{},
				"put:ServicePublicityAction"),
			beego.NSRouter("/services/:id([0-9]+)/autoscale",
				&AutoScaleController{},
				"post:CreateAutoScaleAction;get:ListAutoScaleAction"),
			beego.NSRouter("/services/:id([0-9]+)/autoscale/:hpaid([0-9]+)",
				&AutoScaleController{},
				"put:UpdateAutoScaleAction;delete:DeleteAutoScaleAction"),
			beego.NSRouter("/files/upload",
				&FileUploadController{},
				"post:Upload"),
			beego.NSRouter("/files/download",
				&FileUploadController{},
				"head:DownloadProbe;get:Download"),
			beego.NSRouter("/files/list",
				&FileUploadController{},
				"post:ListFiles"),
			beego.NSRouter("/files/remove",
				&FileUploadController{},
				"post:RemoveFile"),
			beego.NSRouter("/jenkins-job/:userID([0-9]+)/:buildNumber([0-9]+)",
				&JenkinsJobCallbackController{},
				"get:BuildNumberCallback"),
			beego.NSRouter("/jenkins-job/console",
				&JenkinsJobController{},
				"get:Console"),
			beego.NSRouter("/jenkins-job/stop",
				&JenkinsJobController{},
				"get:Stop"),
			beego.NSRouter("/email/ping",
				&EmailController{},
				"post:Ping"),
			beego.NSRouter("/email/notification",
				&EmailController{},
				"post:GrafanaNotification"),
			beego.NSRouter("/operations",
				&OperationController{},
				"get:OperationList;post:CreateOperation"),
			beego.NSRouter("/helm/repositories",
				&HelmController{},
				"get:ListHelmReposAction"),
			beego.NSRouter("/helm/repositories/:id([0-9]+)",
				&HelmController{},
				"get:GetHelmRepoDetailAction"),
			beego.NSRouter("/helm/repositories/:id([0-9]+)/chartupload",
				&HelmController{},
				"post:UploadHelmChartAction"),
			beego.NSRouter("/helm/repositories/:id([0-9]+)/charts/:chartname(.+)/:chartversion(.+)",
				&HelmController{},
				"get:GetHelmChartDetailAction;delete:DeleteHelmChartAction"),
			beego.NSRouter("/helm/release/existing",
				&HelmController{},
				"get:ReleaseExists"),
			beego.NSRouter("/helm/release",
				&HelmController{},
				"get:ListHelmReleaseAction;post:InstallHelmChartAction"),
			beego.NSRouter("/helm/release/:id([0-9]+)",
				&HelmController{},
				"delete:DeleteHelmReleaseAction;get:GetHelmReleaseAction"),
			beego.NSRouter("/helm/release/:id([0-9]+)/logs/:podname(.+)",
				&HelmController{},
				"get:GetReleasePodLogsAction"),
			beego.NSRouter("/jobs",
				&JobController{},
				"get:GetJobListAction"),
			beego.NSRouter("/jobs/deployment",
				&JobController{},
				"post:DeployJobAction"),
			beego.NSRouter("/jobs/:id([0-9]+)",
				&JobController{},
				"delete:DeleteJobAction"),
			beego.NSRouter("/jobs/:id([0-9]+)/status",
				&JobController{},
				"get:GetJobStatusAction"),
			beego.NSRouter("/jobs/:id([0-9]+)/config",
				&JobController{},
				"get:GetJobConfigAction"),
			beego.NSRouter("/jobs/:id([0-9]+)/pods",
				&JobController{},
				"get:GetJobPodsAction"),
			beego.NSRouter("/jobs/:id([0-9]+)/logs/:podname(.+)",
				&JobController{},
				"get:GetJobLogsAction"),
			beego.NSRouter("/jobs/exists",
				&JobController{},
				"get:JobExists"),
			beego.NSRouter("/jobs/selectjobs",
				&JobController{},
				"get:GetSelectableJobsAction"),
			beego.NSRouter("/pods/:projectid([0-9]+)/:podname(.+)/shell",
				&PodController{},
				"get:PodShellAction"),
			beego.NSRouter("/pods/:projectid([0-9]+)/:podname(.+)/download",
				&PodController{},
				"get:CopyFromPodAction"),
			beego.NSRouter("/pods/:projectid([0-9]+)/:podname(.+)/upload",
				&PodController{},
				"post:CopyToPodAction"),
			beego.NSRouter("/forgot-password",
				&EmailController{},
				"post:ForgotPasswordEmail"),
			beego.NSRouter("/k8sproxy",
				&K8SProxyController{},
				"get:GetK8SProxyConfig;put:SetK8SProxyConfig"),
		),
	)

	beego.AddNamespace(ns)
	beego.Router("/deploy/:owner_name/:project_name/:service_name", &ServiceShowController{})
	beego.SetStaticPath("/swagger", "swagger")
	beego.Router("/kubernetes/?:all(.*)", &K8SProxyController{}, "*:ProxyAction")
}
