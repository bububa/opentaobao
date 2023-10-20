package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeBuyAPIRequest 手淘下单能力开放 API请求
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
type AlibabaInteractSensorTradeBuyAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractSensorTradeBuyRequest 初始化AlibabaInteractSensorTradeBuyAPIRequest对象
func NewAlibabaInteractSensorTradeBuyRequest() *AlibabaInteractSensorTradeBuyAPIRequest {
	return &AlibabaInteractSensorTradeBuyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorTradeBuyAPIRequest) Reset() {
	r._id = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade.buy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeBuyAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractSensorTradeBuyAPIRequest) GetId() string {
	return r._id
}

var poolAlibabaInteractSensorTradeBuyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorTradeBuyRequest()
	},
}

// GetAlibabaInteractSensorTradeBuyRequest 从 sync.Pool 获取 AlibabaInteractSensorTradeBuyAPIRequest
func GetAlibabaInteractSensorTradeBuyAPIRequest() *AlibabaInteractSensorTradeBuyAPIRequest {
	return poolAlibabaInteractSensorTradeBuyAPIRequest.Get().(*AlibabaInteractSensorTradeBuyAPIRequest)
}

// ReleaseAlibabaInteractSensorTradeBuyAPIRequest 将 AlibabaInteractSensorTradeBuyAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorTradeBuyAPIRequest(v *AlibabaInteractSensorTradeBuyAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorTradeBuyAPIRequest.Put(v)
}
