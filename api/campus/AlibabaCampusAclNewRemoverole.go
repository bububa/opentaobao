package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewRemoverole 删除角色
// alibaba.campus.acl.new.removerole
//
// 删除角色
func AlibabaCampusAclNewRemoverole(clt *core.SDKClient, req *campus.AlibabaCampusAclNewRemoveroleAPIRequest, resp *campus.AlibabaCampusAclNewRemoveroleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
