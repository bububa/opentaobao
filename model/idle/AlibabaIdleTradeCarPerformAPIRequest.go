package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTradeCarPerformAPIRequest 二手车寄卖履约接口 API请求
// alibaba.idle.trade.car.perform
//
// 二手车寄卖履约接口
type AlibabaIdleTradeCarPerformAPIRequest struct {
	model.Params
	// 履约参数
	_carConsignmentParam *CarConsignmentParam
}

// NewAlibabaIdleTradeCarPerformRequest 初始化AlibabaIdleTradeCarPerformAPIRequest对象
func NewAlibabaIdleTradeCarPerformRequest() *AlibabaIdleTradeCarPerformAPIRequest {
	return &AlibabaIdleTradeCarPerformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTradeCarPerformAPIRequest) Reset() {
	r._carConsignmentParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTradeCarPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.trade.car.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTradeCarPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTradeCarPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCarConsignmentParam is CarConsignmentParam Setter
// 履约参数
func (r *AlibabaIdleTradeCarPerformAPIRequest) SetCarConsignmentParam(_carConsignmentParam *CarConsignmentParam) error {
	r._carConsignmentParam = _carConsignmentParam
	r.Set("car_consignment_param", _carConsignmentParam)
	return nil
}

// GetCarConsignmentParam CarConsignmentParam Getter
func (r AlibabaIdleTradeCarPerformAPIRequest) GetCarConsignmentParam() *CarConsignmentParam {
	return r._carConsignmentParam
}

var poolAlibabaIdleTradeCarPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTradeCarPerformRequest()
	},
}

// GetAlibabaIdleTradeCarPerformRequest 从 sync.Pool 获取 AlibabaIdleTradeCarPerformAPIRequest
func GetAlibabaIdleTradeCarPerformAPIRequest() *AlibabaIdleTradeCarPerformAPIRequest {
	return poolAlibabaIdleTradeCarPerformAPIRequest.Get().(*AlibabaIdleTradeCarPerformAPIRequest)
}

// ReleaseAlibabaIdleTradeCarPerformAPIRequest 将 AlibabaIdleTradeCarPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTradeCarPerformAPIRequest(v *AlibabaIdleTradeCarPerformAPIRequest) {
	v.Reset()
	poolAlibabaIdleTradeCarPerformAPIRequest.Put(v)
}
