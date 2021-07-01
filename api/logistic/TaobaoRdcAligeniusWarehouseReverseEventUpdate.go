package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoRdcAligeniusWarehouseReverseEventUpdate
销退单事件回传接口
taobao.rdc.aligenius.warehouse.reverse.event.update

用于erp回传销退单相关信息到平台 */
func TaobaoRdcAligeniusWarehouseReverseEventUpdate(clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest, session string) (*logistic.TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse, error) {
	var resp logistic.TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
