package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclcancelrolesfromuser 撤销用户授予的角色
// alibaba.campus.acl.cancelrolesfromuser
//
// 撤销用户授予的角色
func Alibabacampusaclcancelrolesfromuser(clt *core.SDKClient, req *campus.AlibabacampusaclcancelrolesfromuserAPIRequest, session string) (*campus.AlibabacampusaclcancelrolesfromuserAPIResponse, error) {
	var resp campus.AlibabacampusaclcancelrolesfromuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
