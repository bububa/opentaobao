package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGetpermissionbyroleid 根据角色Id查询权限
// alibaba.campus.acl.getpermissionbyroleid
//
// 根据角色查询权限
func AlibabaCampusAclGetpermissionbyroleid(clt *core.SDKClient, req *campus.AlibabaCampusAclGetpermissionbyroleidAPIRequest, session string) (*campus.AlibabaCampusAclGetpermissionbyroleidAPIResponse, error) {
	var resp campus.AlibabaCampusAclGetpermissionbyroleidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
