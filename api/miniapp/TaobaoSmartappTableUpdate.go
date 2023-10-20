package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaosmartapptableupdate 智能应用服务登记工作表数据更新
// taobao.smartapp.table.update
//
// 智能应用服务登记工作表数据更新
func Taobaosmartapptableupdate(clt *core.SDKClient, req *miniapp.TaobaosmartapptableupdateAPIRequest, session string) (*miniapp.TaobaosmartapptableupdateAPIResponse, error) {
	var resp miniapp.TaobaosmartapptableupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
