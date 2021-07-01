package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstShiporderCancelAPIRequest
零售通发货单取消 API请求
alibaba.lst.shiporder.cancel

通过该接口可以取消零售通运保保发货单，并处理相关业务流程。 */
type AlibabaLstShiporderCancelAPIRequest struct {
	model.Params
	// 取消原因
	_reason string
	// 订单号
	_outOrderId string
	// 需要退款的明细ID
	_detailOrderIds []string
}

// NewAlibabaLstShiporderCancelRequest 初始化AlibabaLstShiporderCancelAPIRequest对象
func NewAlibabaLstShiporderCancelRequest() *AlibabaLstShiporderCancelAPIRequest {
	return &AlibabaLstShiporderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstShiporderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstShiporderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Reason Setter
// 取消原因
func (r *AlibabaLstShiporderCancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r AlibabaLstShiporderCancelAPIRequest) GetReason() string {
	return r._reason
}

// Set is OutOrderId Setter
// 订单号
func (r *AlibabaLstShiporderCancelAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// Get OutOrderId Getter
func (r AlibabaLstShiporderCancelAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// Set is DetailOrderIds Setter
// 需要退款的明细ID
func (r *AlibabaLstShiporderCancelAPIRequest) SetDetailOrderIds(_detailOrderIds []string) error {
	r._detailOrderIds = _detailOrderIds
	r.Set("detail_order_ids", _detailOrderIds)
	return nil
}

// Get DetailOrderIds Getter
func (r AlibabaLstShiporderCancelAPIRequest) GetDetailOrderIds() []string {
	return r._detailOrderIds
}
