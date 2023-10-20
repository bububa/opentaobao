package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGetrolebyempid 根据用户查询角色
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
func AlibabaCampusAclGetrolebyempid(clt *core.SDKClient, req *campus.AlibabaCampusAclGetrolebyempidAPIRequest, resp *campus.AlibabaCampusAclGetrolebyempidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
