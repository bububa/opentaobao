package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsWmsOrderMaterialSync 仓服务商订单包材耗材信息同步
// taobao.logistics.wms.order.material.sync
//
// 仓服务商订单包材耗材信息同步
func TaobaoLogisticsWmsOrderMaterialSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsWmsOrderMaterialSyncAPIRequest, session string) (*logistic.TaobaoLogisticsWmsOrderMaterialSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsWmsOrderMaterialSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
