package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewGetrolewithmenutreenodes 根据角色id查询权限
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
func AlibabaCampusAclNewGetrolewithmenutreenodes(clt *core.SDKClient, req *campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest, resp *campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
