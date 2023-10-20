package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomWttOpentradeCreateorderAPIRequest 充值送活动下单接口 API请求
// alibaba.alicom.wtt.opentrade.createorder
//
// 提供给话费宝创建淘宝订单
type AlibabaAlicomWttOpentradeCreateorderAPIRequest struct {
	model.Params
	// 入参请求说明
	_param0 *OpentradCreateOrderRequestDto
}

// NewAlibabaAlicomWttOpentradeCreateorderRequest 初始化AlibabaAlicomWttOpentradeCreateorderAPIRequest对象
func NewAlibabaAlicomWttOpentradeCreateorderRequest() *AlibabaAlicomWttOpentradeCreateorderAPIRequest {
	return &AlibabaAlicomWttOpentradeCreateorderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomWttOpentradeCreateorderAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeCreateorderAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeCreateorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomWttOpentradeCreateorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参请求说明
func (r *AlibabaAlicomWttOpentradeCreateorderAPIRequest) SetParam0(_param0 *OpentradCreateOrderRequestDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAlicomWttOpentradeCreateorderAPIRequest) GetParam0() *OpentradCreateOrderRequestDto {
	return r._param0
}

var poolAlibabaAlicomWttOpentradeCreateorderAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomWttOpentradeCreateorderRequest()
	},
}

// GetAlibabaAlicomWttOpentradeCreateorderRequest 从 sync.Pool 获取 AlibabaAlicomWttOpentradeCreateorderAPIRequest
func GetAlibabaAlicomWttOpentradeCreateorderAPIRequest() *AlibabaAlicomWttOpentradeCreateorderAPIRequest {
	return poolAlibabaAlicomWttOpentradeCreateorderAPIRequest.Get().(*AlibabaAlicomWttOpentradeCreateorderAPIRequest)
}

// ReleaseAlibabaAlicomWttOpentradeCreateorderAPIRequest 将 AlibabaAlicomWttOpentradeCreateorderAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomWttOpentradeCreateorderAPIRequest(v *AlibabaAlicomWttOpentradeCreateorderAPIRequest) {
	v.Reset()
	poolAlibabaAlicomWttOpentradeCreateorderAPIRequest.Put(v)
}
