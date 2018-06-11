export enum MESSAGE_TARGET {
  TOGGLE_PROJECT, DELETE_PROJECT, TOGGLE_NODE, DELETE_SERVICE,
  TOGGLE_SERVICE, DELETE_IMAGE, CANCEL_BUILD_IMAGE, CANCEL_BUILD_SERVICE,
  CANCEL_BUILD_SERVICE_GUARD, FORCE_QUIT_BUILD_IMAGE, DELETE_TAG, DELETE_NODE_GROUP,
  DELETE_SERVICE_DEPLOYMENT,DELETE_USER,SIGN_IN_ERROR
}

export const DISMISS_INLINE_ALERT_INTERVAL: number = 4 * 1000;
export const DISMISS_GLOBAL_ALERT_INTERVAL: number = 10 * 1000;

export enum MESSAGE_TYPE {
  NONE, COMMON_ERROR = 1, INVALID_USER, INTERNAL_ERROR, SHOW_DETAIL
}

export const ROLES: {[key: number]: string} = {
  1: 'PROJECT.PROJECT_ADMIN', 2: 'PROJECT.DEVELOPER', 3: 'PROJECT.VISITOR'
};

export enum BUTTON_STYLE {
  CONFIRMATION = 1, DELETION, YES_NO, ONLY_CONFIRM
}

export enum SERVICE_STATUS{
  PREPARING,
  RUNNING,
  STOPPED,
  WARNING
}

export enum GUIDE_STEP{
  NONE_STEP,
  PROJECT_LIST,
  CREATE_PROJECT,
  SERVICE_LIST,
  CREATE_SERVICE
}