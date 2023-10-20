package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptablefulldataget 智能应用工作表地址查询
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
func Taobaosmartapptablefulldataget(clt *core.SDKClient, req *miniapp.TaobaosmartapptablefulldatagetAPIRequest, session string) (*miniapp.TaobaosmartapptablefulldatagetAPIResponse, error) {
	var resp miniapp.TaobaosmartapptablefulldatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
