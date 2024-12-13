package user

import (
	"github.com/openimsdk/protocol/rpccall"
	"google.golang.org/grpc"
)

func InitUser(conn *grpc.ClientConn) {
	GetDesignateUsersCaller.SetConn(conn)
	UpdateUserInfoCaller.SetConn(conn)
	UpdateUserInfoExCaller.SetConn(conn)
	SetGlobalRecvMessageOptCaller.SetConn(conn)
	GetGlobalRecvMessageOptCaller.SetConn(conn)
	AccountCheckCaller.SetConn(conn)
	GetPaginationUsersCaller.SetConn(conn)
	UserRegisterCaller.SetConn(conn)
	GetAllUserIDCaller.SetConn(conn)
	UserRegisterCountCaller.SetConn(conn)
	SubscribeOrCancelUsersStatusCaller.SetConn(conn)
	GetSubscribeUsersStatusCaller.SetConn(conn)
	GetUserStatusCaller.SetConn(conn)
	SetUserStatusCaller.SetConn(conn)
	ProcessUserCommandAddCaller.SetConn(conn)
	ProcessUserCommandUpdateCaller.SetConn(conn)
	ProcessUserCommandDeleteCaller.SetConn(conn)
	ProcessUserCommandGetCaller.SetConn(conn)
	ProcessUserCommandGetAllCaller.SetConn(conn)
	AddNotificationAccountCaller.SetConn(conn)
	UpdateNotificationAccountInfoCaller.SetConn(conn)
	SearchNotificationAccountCaller.SetConn(conn)
	GetNotificationAccountCaller.SetConn(conn)
	SortQueryCaller.SetConn(conn)
	SetUserOnlineStatusCaller.SetConn(conn)
	GetAllOnlineUsersCaller.SetConn(conn)
}

var (
	GetDesignateUsersCaller             = rpccall.NewRpcCaller[GetDesignateUsersReq, GetDesignateUsersResp](User_GetDesignateUsers_FullMethodName)
	UpdateUserInfoCaller                = rpccall.NewRpcCaller[UpdateUserInfoReq, UpdateUserInfoResp](User_UpdateUserInfo_FullMethodName)
	UpdateUserInfoExCaller              = rpccall.NewRpcCaller[UpdateUserInfoExReq, UpdateUserInfoExResp](User_UpdateUserInfoEx_FullMethodName)
	SetGlobalRecvMessageOptCaller       = rpccall.NewRpcCaller[SetGlobalRecvMessageOptReq, SetGlobalRecvMessageOptResp](User_SetGlobalRecvMessageOpt_FullMethodName)
	GetGlobalRecvMessageOptCaller       = rpccall.NewRpcCaller[GetGlobalRecvMessageOptReq, GetGlobalRecvMessageOptResp](User_GetGlobalRecvMessageOpt_FullMethodName)
	AccountCheckCaller                  = rpccall.NewRpcCaller[AccountCheckReq, AccountCheckResp](User_AccountCheck_FullMethodName)
	GetPaginationUsersCaller            = rpccall.NewRpcCaller[GetPaginationUsersReq, GetPaginationUsersResp](User_GetPaginationUsers_FullMethodName)
	UserRegisterCaller                  = rpccall.NewRpcCaller[UserRegisterReq, UserRegisterResp](User_UserRegister_FullMethodName)
	GetAllUserIDCaller                  = rpccall.NewRpcCaller[GetAllUserIDReq, GetAllUserIDResp](User_GetAllUserID_FullMethodName)
	UserRegisterCountCaller             = rpccall.NewRpcCaller[UserRegisterCountReq, UserRegisterCountResp](User_UserRegisterCount_FullMethodName)
	SubscribeOrCancelUsersStatusCaller  = rpccall.NewRpcCaller[SubscribeOrCancelUsersStatusReq, SubscribeOrCancelUsersStatusResp](User_SubscribeOrCancelUsersStatus_FullMethodName)
	GetSubscribeUsersStatusCaller       = rpccall.NewRpcCaller[GetSubscribeUsersStatusReq, GetSubscribeUsersStatusResp](User_GetSubscribeUsersStatus_FullMethodName)
	GetUserStatusCaller                 = rpccall.NewRpcCaller[GetUserStatusReq, GetUserStatusResp](User_GetUserStatus_FullMethodName)
	SetUserStatusCaller                 = rpccall.NewRpcCaller[SetUserStatusReq, SetUserStatusResp](User_SetUserStatus_FullMethodName)
	ProcessUserCommandAddCaller         = rpccall.NewRpcCaller[ProcessUserCommandAddReq, ProcessUserCommandAddResp](User_ProcessUserCommandAdd_FullMethodName)
	ProcessUserCommandUpdateCaller      = rpccall.NewRpcCaller[ProcessUserCommandUpdateReq, ProcessUserCommandUpdateResp](User_ProcessUserCommandUpdate_FullMethodName)
	ProcessUserCommandDeleteCaller      = rpccall.NewRpcCaller[ProcessUserCommandDeleteReq, ProcessUserCommandDeleteResp](User_ProcessUserCommandDelete_FullMethodName)
	ProcessUserCommandGetCaller         = rpccall.NewRpcCaller[ProcessUserCommandGetReq, ProcessUserCommandGetResp](User_ProcessUserCommandGet_FullMethodName)
	ProcessUserCommandGetAllCaller      = rpccall.NewRpcCaller[ProcessUserCommandGetAllReq, ProcessUserCommandGetAllResp](User_ProcessUserCommandGetAll_FullMethodName)
	AddNotificationAccountCaller        = rpccall.NewRpcCaller[AddNotificationAccountReq, AddNotificationAccountResp](User_AddNotificationAccount_FullMethodName)
	UpdateNotificationAccountInfoCaller = rpccall.NewRpcCaller[UpdateNotificationAccountInfoReq, UpdateNotificationAccountInfoResp](User_UpdateNotificationAccountInfo_FullMethodName)
	SearchNotificationAccountCaller     = rpccall.NewRpcCaller[SearchNotificationAccountReq, SearchNotificationAccountResp](User_SearchNotificationAccount_FullMethodName)
	GetNotificationAccountCaller        = rpccall.NewRpcCaller[GetNotificationAccountReq, GetNotificationAccountResp](User_GetNotificationAccount_FullMethodName)
	SortQueryCaller                     = rpccall.NewRpcCaller[SortQueryReq, SortQueryResp](User_SortQuery_FullMethodName)
	SetUserOnlineStatusCaller           = rpccall.NewRpcCaller[SetUserOnlineStatusReq, SetUserOnlineStatusResp](User_SetUserOnlineStatus_FullMethodName)
	GetAllOnlineUsersCaller             = rpccall.NewRpcCaller[GetAllOnlineUsersReq, GetAllOnlineUsersResp](User_GetAllOnlineUsers_FullMethodName)
)
