package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckuserpermission 校验用户是否有权限
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
func AlibabaCampusAclNewCheckuserpermission(clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserpermissionAPIRequest, resp *campus.AlibabaCampusAclNewCheckuserpermissionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
