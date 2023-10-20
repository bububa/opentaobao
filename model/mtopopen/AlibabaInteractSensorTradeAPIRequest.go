package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeAPIRequest 交易组件 API请求
// alibaba.interact.sensor.trade
//
// 交易流程
type AlibabaInteractSensorTradeAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractSensorTradeRequest 初始化AlibabaInteractSensorTradeAPIRequest对象
func NewAlibabaInteractSensorTradeRequest() *AlibabaInteractSensorTradeAPIRequest {
	return &AlibabaInteractSensorTradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractSensorTradeAPIRequest) Reset() {
	r._id = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTradeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorTradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorTradeAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractSensorTradeAPIRequest) GetId() string {
	return r._id
}

var poolAlibabaInteractSensorTradeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractSensorTradeRequest()
	},
}

// GetAlibabaInteractSensorTradeRequest 从 sync.Pool 获取 AlibabaInteractSensorTradeAPIRequest
func GetAlibabaInteractSensorTradeAPIRequest() *AlibabaInteractSensorTradeAPIRequest {
	return poolAlibabaInteractSensorTradeAPIRequest.Get().(*AlibabaInteractSensorTradeAPIRequest)
}

// ReleaseAlibabaInteractSensorTradeAPIRequest 将 AlibabaInteractSensorTradeAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractSensorTradeAPIRequest(v *AlibabaInteractSensorTradeAPIRequest) {
	v.Reset()
	poolAlibabaInteractSensorTradeAPIRequest.Put(v)
}
