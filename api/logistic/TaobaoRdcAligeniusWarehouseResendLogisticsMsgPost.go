package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPost 补发单物流信息回传
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
func TaobaoRdcAligeniusWarehouseResendLogisticsMsgPost(clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest, session string) (*logistic.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse, error) {
	var resp logistic.TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
