package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardordersStatusMakeFinishAPIRequest 制卡商通知制卡完成 API请求
// alibaba.fundplatform.cardorders.status.make.finish
//
// 当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
type AlibabaFundplatformCardordersStatusMakeFinishAPIRequest struct {
	model.Params
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabaFundplatformCardordersStatusMakeFinishRequest 初始化AlibabaFundplatformCardordersStatusMakeFinishAPIRequest对象
func NewAlibabaFundplatformCardordersStatusMakeFinishRequest() *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest {
	return &AlibabaFundplatformCardordersStatusMakeFinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.status.make.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// Get CardOrderId Getter
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
