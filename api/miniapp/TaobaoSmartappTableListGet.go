package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptablelistget 智能应用服务登记工作表列表查询
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
func Taobaosmartapptablelistget(clt *core.SDKClient, req *miniapp.TaobaosmartapptablelistgetAPIRequest, session string) (*miniapp.TaobaosmartapptablelistgetAPIResponse, error) {
	var resp miniapp.TaobaosmartapptablelistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
