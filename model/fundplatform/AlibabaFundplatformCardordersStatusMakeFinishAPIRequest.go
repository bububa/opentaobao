package fundplatform

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) Reset() {
	r._cardOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorders.status.make.finish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}

var poolAlibabaFundplatformCardordersStatusMakeFinishAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardordersStatusMakeFinishRequest()
	},
}

// GetAlibabaFundplatformCardordersStatusMakeFinishRequest 从 sync.Pool 获取 AlibabaFundplatformCardordersStatusMakeFinishAPIRequest
func GetAlibabaFundplatformCardordersStatusMakeFinishAPIRequest() *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest {
	return poolAlibabaFundplatformCardordersStatusMakeFinishAPIRequest.Get().(*AlibabaFundplatformCardordersStatusMakeFinishAPIRequest)
}

// ReleaseAlibabaFundplatformCardordersStatusMakeFinishAPIRequest 将 AlibabaFundplatformCardordersStatusMakeFinishAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardordersStatusMakeFinishAPIRequest(v *AlibabaFundplatformCardordersStatusMakeFinishAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardordersStatusMakeFinishAPIRequest.Put(v)
}
