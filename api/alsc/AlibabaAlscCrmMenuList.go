package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmmenulist 获取特价菜单
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
func Alibabaalsccrmmenulist(clt *core.SDKClient, req *alsc.AlibabaalsccrmmenulistAPIRequest, session string) (*alsc.AlibabaalsccrmmenulistAPIResponse, error) {
	var resp alsc.AlibabaalsccrmmenulistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
