package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMenuList 获取特价菜单
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
func AlibabaAlscCrmMenuList(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMenuListAPIRequest, session string) (*alsc.AlibabaAlscCrmMenuListAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmMenuListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
