package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGrantpermiitemtorole 权限赋予角色
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
func AlibabaCampusAclGrantpermiitemtorole(clt *core.SDKClient, req *campus.AlibabaCampusAclGrantpermiitemtoroleAPIRequest, session string) (*campus.AlibabaCampusAclGrantpermiitemtoroleAPIResponse, error) {
	var resp campus.AlibabaCampusAclGrantpermiitemtoroleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
