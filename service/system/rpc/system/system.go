// Code generated by goctl. DO NOT EDIT!
// Source: system.proto

package system

import (
	"context"

	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq          = systemclient.ConfigAddReq
	ConfigAddResp         = systemclient.ConfigAddResp
	ConfigDeleteReq       = systemclient.ConfigDeleteReq
	ConfigDeleteResp      = systemclient.ConfigDeleteResp
	ConfigListData        = systemclient.ConfigListData
	ConfigListReq         = systemclient.ConfigListReq
	ConfigListResp        = systemclient.ConfigListResp
	ConfigUpdateReq       = systemclient.ConfigUpdateReq
	ConfigUpdateResp      = systemclient.ConfigUpdateResp
	DeptAddReq            = systemclient.DeptAddReq
	DeptAddResp           = systemclient.DeptAddResp
	DeptDeleteReq         = systemclient.DeptDeleteReq
	DeptDeleteResp        = systemclient.DeptDeleteResp
	DeptListData          = systemclient.DeptListData
	DeptListReq           = systemclient.DeptListReq
	DeptListResp          = systemclient.DeptListResp
	DeptUpdateReq         = systemclient.DeptUpdateReq
	DeptUpdateResp        = systemclient.DeptUpdateResp
	DictAddReq            = systemclient.DictAddReq
	DictAddResp           = systemclient.DictAddResp
	DictDeleteReq         = systemclient.DictDeleteReq
	DictDeleteResp        = systemclient.DictDeleteResp
	DictListData          = systemclient.DictListData
	DictListReq           = systemclient.DictListReq
	DictListResp          = systemclient.DictListResp
	DictUpdateReq         = systemclient.DictUpdateReq
	DictUpdateResp        = systemclient.DictUpdateResp
	InfoReq               = systemclient.InfoReq
	InfoResp              = systemclient.InfoResp
	JobAddReq             = systemclient.JobAddReq
	JobAddResp            = systemclient.JobAddResp
	JobDeleteReq          = systemclient.JobDeleteReq
	JobDeleteResp         = systemclient.JobDeleteResp
	JobListData           = systemclient.JobListData
	JobListReq            = systemclient.JobListReq
	JobListResp           = systemclient.JobListResp
	JobUpdateReq          = systemclient.JobUpdateReq
	JobUpdateResp         = systemclient.JobUpdateResp
	LoginLogAddReq        = systemclient.LoginLogAddReq
	LoginLogAddResp       = systemclient.LoginLogAddResp
	LoginLogDeleteReq     = systemclient.LoginLogDeleteReq
	LoginLogDeleteResp    = systemclient.LoginLogDeleteResp
	LoginLogListData      = systemclient.LoginLogListData
	LoginLogListReq       = systemclient.LoginLogListReq
	LoginLogListResp      = systemclient.LoginLogListResp
	LoginReq              = systemclient.LoginReq
	LoginResp             = systemclient.LoginResp
	MenuAddReq            = systemclient.MenuAddReq
	MenuAddResp           = systemclient.MenuAddResp
	MenuDeleteReq         = systemclient.MenuDeleteReq
	MenuDeleteResp        = systemclient.MenuDeleteResp
	MenuListData          = systemclient.MenuListData
	MenuListReq           = systemclient.MenuListReq
	MenuListResp          = systemclient.MenuListResp
	MenuListTree          = systemclient.MenuListTree
	MenuUpdateReq         = systemclient.MenuUpdateReq
	MenuUpdateResp        = systemclient.MenuUpdateResp
	QueryMenuByRoleIdReq  = systemclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = systemclient.QueryMenuByRoleIdResp
	ReSetPasswordReq      = systemclient.ReSetPasswordReq
	ReSetPasswordResp     = systemclient.ReSetPasswordResp
	RoleAddReq            = systemclient.RoleAddReq
	RoleAddResp           = systemclient.RoleAddResp
	RoleDeleteReq         = systemclient.RoleDeleteReq
	RoleDeleteResp        = systemclient.RoleDeleteResp
	RoleListData          = systemclient.RoleListData
	RoleListReq           = systemclient.RoleListReq
	RoleListResp          = systemclient.RoleListResp
	RoleUpdateReq         = systemclient.RoleUpdateReq
	RoleUpdateResp        = systemclient.RoleUpdateResp
	SystemLogAddReq       = systemclient.SystemLogAddReq
	SystemLogAddResp      = systemclient.SystemLogAddResp
	SystemLogDeleteReq    = systemclient.SystemLogDeleteReq
	SystemLogDeleteResp   = systemclient.SystemLogDeleteResp
	SystemLogListData     = systemclient.SystemLogListData
	SystemLogListReq      = systemclient.SystemLogListReq
	SystemLogListResp     = systemclient.SystemLogListResp
	UpdateConfigRoleReq   = systemclient.UpdateConfigRoleReq
	UpdateConfigRoleResp  = systemclient.UpdateConfigRoleResp
	UpdateMenuRoleReq     = systemclient.UpdateMenuRoleReq
	UpdateMenuRoleResp    = systemclient.UpdateMenuRoleResp
	UpdateRoleRoleReq     = systemclient.UpdateRoleRoleReq
	UpdateRoleRoleResp    = systemclient.UpdateRoleRoleResp
	UserAddReq            = systemclient.UserAddReq
	UserAddResp           = systemclient.UserAddResp
	UserDeleteReq         = systemclient.UserDeleteReq
	UserDeleteResp        = systemclient.UserDeleteResp
	UserListData          = systemclient.UserListData
	UserListReq           = systemclient.UserListReq
	UserListResp          = systemclient.UserListResp
	UserStatusReq         = systemclient.UserStatusReq
	UserStatusResp        = systemclient.UserStatusResp
	UserUpdateReq         = systemclient.UserUpdateReq
	UserUpdateResp        = systemclient.UserUpdateResp

	System interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error)
		UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq, opts ...grpc.CallOption) (*UpdateRoleRoleResp, error)
		QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error)
		UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
		DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error)
		DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
		DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error)
		DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error)
		DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error)
		DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
		DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error)
		DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error)
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		SystemLogAdd(ctx context.Context, in *SystemLogAddReq, opts ...grpc.CallOption) (*SystemLogAddResp, error)
		SystemLogList(ctx context.Context, in *SystemLogListReq, opts ...grpc.CallOption) (*SystemLogListResp, error)
		SystemLogDelete(ctx context.Context, in *SystemLogDeleteReq, opts ...grpc.CallOption) (*SystemLogDeleteResp, error)
		ConfigAdd(ctx context.Context, in *ConfigAddReq, opts ...grpc.CallOption) (*ConfigAddResp, error)
		ConfigList(ctx context.Context, in *ConfigListReq, opts ...grpc.CallOption) (*ConfigListResp, error)
		ConfigUpdate(ctx context.Context, in *ConfigUpdateReq, opts ...grpc.CallOption) (*ConfigUpdateResp, error)
		ConfigDelete(ctx context.Context, in *ConfigDeleteReq, opts ...grpc.CallOption) (*ConfigDeleteResp, error)
		UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq, opts ...grpc.CallOption) (*UpdateConfigRoleResp, error)
		JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error)
		JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error)
		JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error)
		JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error)
	}

	defaultSystem struct {
		cli zrpc.Client
	}
)

func NewSystem(cli zrpc.Client) System {
	return &defaultSystem{
		cli: cli,
	}
}

func (m *defaultSystem) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSystem) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSystem) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSystem) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSystem) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultSystem) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultSystem) ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.ReSetPassword(ctx, in, opts...)
}

func (m *defaultSystem) UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

func (m *defaultSystem) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultSystem) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultSystem) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultSystem) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (m *defaultSystem) UpdateRoleRole(ctx context.Context, in *UpdateRoleRoleReq, opts ...grpc.CallOption) (*UpdateRoleRoleResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UpdateRoleRole(ctx, in, opts...)
}

func (m *defaultSystem) QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.QueryMenuByRoleId(ctx, in, opts...)
}

func (m *defaultSystem) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UpdateMenuRole(ctx, in, opts...)
}

func (m *defaultSystem) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultSystem) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultSystem) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultSystem) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

func (m *defaultSystem) DictAdd(ctx context.Context, in *DictAddReq, opts ...grpc.CallOption) (*DictAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DictAdd(ctx, in, opts...)
}

func (m *defaultSystem) DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DictList(ctx, in, opts...)
}

func (m *defaultSystem) DictUpdate(ctx context.Context, in *DictUpdateReq, opts ...grpc.CallOption) (*DictUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DictUpdate(ctx, in, opts...)
}

func (m *defaultSystem) DictDelete(ctx context.Context, in *DictDeleteReq, opts ...grpc.CallOption) (*DictDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DictDelete(ctx, in, opts...)
}

func (m *defaultSystem) DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DeptAdd(ctx, in, opts...)
}

func (m *defaultSystem) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DeptList(ctx, in, opts...)
}

func (m *defaultSystem) DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in, opts...)
}

func (m *defaultSystem) DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.DeptDelete(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

func (m *defaultSystem) SystemLogAdd(ctx context.Context, in *SystemLogAddReq, opts ...grpc.CallOption) (*SystemLogAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.SystemLogAdd(ctx, in, opts...)
}

func (m *defaultSystem) SystemLogList(ctx context.Context, in *SystemLogListReq, opts ...grpc.CallOption) (*SystemLogListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.SystemLogList(ctx, in, opts...)
}

func (m *defaultSystem) SystemLogDelete(ctx context.Context, in *SystemLogDeleteReq, opts ...grpc.CallOption) (*SystemLogDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.SystemLogDelete(ctx, in, opts...)
}

func (m *defaultSystem) ConfigAdd(ctx context.Context, in *ConfigAddReq, opts ...grpc.CallOption) (*ConfigAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.ConfigAdd(ctx, in, opts...)
}

func (m *defaultSystem) ConfigList(ctx context.Context, in *ConfigListReq, opts ...grpc.CallOption) (*ConfigListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.ConfigList(ctx, in, opts...)
}

func (m *defaultSystem) ConfigUpdate(ctx context.Context, in *ConfigUpdateReq, opts ...grpc.CallOption) (*ConfigUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.ConfigUpdate(ctx, in, opts...)
}

func (m *defaultSystem) ConfigDelete(ctx context.Context, in *ConfigDeleteReq, opts ...grpc.CallOption) (*ConfigDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.ConfigDelete(ctx, in, opts...)
}

func (m *defaultSystem) UpdateConfigRole(ctx context.Context, in *UpdateConfigRoleReq, opts ...grpc.CallOption) (*UpdateConfigRoleResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.UpdateConfigRole(ctx, in, opts...)
}

func (m *defaultSystem) JobAdd(ctx context.Context, in *JobAddReq, opts ...grpc.CallOption) (*JobAddResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.JobAdd(ctx, in, opts...)
}

func (m *defaultSystem) JobList(ctx context.Context, in *JobListReq, opts ...grpc.CallOption) (*JobListResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.JobList(ctx, in, opts...)
}

func (m *defaultSystem) JobUpdate(ctx context.Context, in *JobUpdateReq, opts ...grpc.CallOption) (*JobUpdateResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.JobUpdate(ctx, in, opts...)
}

func (m *defaultSystem) JobDelete(ctx context.Context, in *JobDeleteReq, opts ...grpc.CallOption) (*JobDeleteResp, error) {
	client := systemclient.NewSystemClient(m.cli.Conn())
	return client.JobDelete(ctx, in, opts...)
}
