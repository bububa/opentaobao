package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseResendUpdate 补发单状态回传
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
func TaobaoRdcAligeniusWarehouseResendUpdate(clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseResendUpdateAPIRequest, session string) (*logistic.TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse, error) {
	var resp logistic.TaobaoRdcAligeniusWarehouseResendUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
