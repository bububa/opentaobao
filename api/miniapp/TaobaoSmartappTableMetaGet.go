package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptablemetaget 智能应用服务登记工作表元数据查询
// taobao.smartapp.table.meta.get
//
// 智能应用服务登记工作表元数据查询
func Taobaosmartapptablemetaget(clt *core.SDKClient, req *miniapp.TaobaosmartapptablemetagetAPIRequest, session string) (*miniapp.TaobaosmartapptablemetagetAPIResponse, error) {
	var resp miniapp.TaobaosmartapptablemetagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
