package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoNextoneLogisticsWarehouseUpdate AG退货入仓状态写接口
// taobao.nextone.logistics.warehouse.update
//
// 商家上传退货入仓状态给ag
func TaobaoNextoneLogisticsWarehouseUpdate(clt *core.SDKClient, req *logistic.TaobaoNextoneLogisticsWarehouseUpdateAPIRequest, resp *logistic.TaobaoNextoneLogisticsWarehouseUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
