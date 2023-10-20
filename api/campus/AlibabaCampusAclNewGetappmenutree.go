package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclNewGetappmenutree 查询应用下的菜单树
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
func AlibabaCampusAclNewGetappmenutree(clt *core.SDKClient, req *campus.AlibabaCampusAclNewGetappmenutreeAPIRequest, session string) (*campus.AlibabaCampusAclNewGetappmenutreeAPIResponse, error) {
	var resp campus.AlibabaCampusAclNewGetappmenutreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
