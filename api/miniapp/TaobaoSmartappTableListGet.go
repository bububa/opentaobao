package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableListGet 智能应用服务登记工作表列表查询
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
func TaobaoSmartappTableListGet(clt *core.SDKClient, req *miniapp.TaobaoSmartappTableListGetAPIRequest, session string) (*miniapp.TaobaoSmartappTableListGetAPIResponse, error) {
	var resp miniapp.TaobaoSmartappTableListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
