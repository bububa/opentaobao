package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordersstatusmakefinishAPIRequest 制卡商通知制卡完成 API请求
// alibaba.fundplatform.cardorders.status.make.finish
//
// 当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
type AlibabafundplatformcardordersstatusmakefinishAPIRequest struct {
	model.Params
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabafundplatformcardordersstatusmakefinishRequest 初始化AlibabafundplatformcardordersstatusmakefinishAPIRequest对象
func NewAlibabafundplatformcardordersstatusmakefinishRequest() *AlibabafundplatformcardordersstatusmakefinishAPIRequest {
	return &AlibabafundplatformcardordersstatusmakefinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordersstatusmakefinishAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.status.make.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordersstatusmakefinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordersstatusmakefinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabafundplatformcardordersstatusmakefinishAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabafundplatformcardordersstatusmakefinishAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
