package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

/* AlibabaCampusAclCheckemprole
校验用户是否有该角色
alibaba.campus.acl.checkemprole

校验用户是否有该权限 */
func AlibabaCampusAclCheckemprole(clt *core.SDKClient, req *campus.AlibabaCampusAclCheckemproleAPIRequest, session string) (*campus.AlibabaCampusAclCheckemproleAPIResponse, error) {
	var resp campus.AlibabaCampusAclCheckemproleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
