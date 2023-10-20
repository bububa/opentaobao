package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptableget 智能应用服务登记工作表数据查询
// taobao.smartapp.table.get
//
// 智能应用服务登记工作表数据查询
func Taobaosmartapptableget(clt *core.SDKClient, req *miniapp.TaobaosmartapptablegetAPIRequest, session string) (*miniapp.TaobaosmartapptablegetAPIResponse, error) {
	var resp miniapp.TaobaosmartapptablegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
