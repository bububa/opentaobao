package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptableadd 智能应用服务登记工作表数据新增
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
func Taobaosmartapptableadd(clt *core.SDKClient, req *miniapp.TaobaosmartapptableaddAPIRequest, session string) (*miniapp.TaobaosmartapptableaddAPIResponse, error) {
	var resp miniapp.TaobaosmartapptableaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
