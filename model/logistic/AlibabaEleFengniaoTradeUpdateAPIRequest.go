package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoTradeUpdateAPIRequest 更新蜂鸟扣费状态 API请求
// alibaba.ele.fengniao.trade.update
//
// 汇金扣费成功后，回调该接口更新扣费状态
type AlibabaEleFengniaoTradeUpdateAPIRequest struct {
	model.Params
	// param 参数
	_param *Param
}

// NewAlibabaEleFengniaoTradeUpdateRequest 初始化AlibabaEleFengniaoTradeUpdateAPIRequest对象
func NewAlibabaEleFengniaoTradeUpdateRequest() *AlibabaEleFengniaoTradeUpdateAPIRequest {
	return &AlibabaEleFengniaoTradeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoTradeUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoTradeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.trade.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoTradeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoTradeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// param 参数
func (r *AlibabaEleFengniaoTradeUpdateAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoTradeUpdateAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoTradeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoTradeUpdateRequest()
	},
}

// GetAlibabaEleFengniaoTradeUpdateRequest 从 sync.Pool 获取 AlibabaEleFengniaoTradeUpdateAPIRequest
func GetAlibabaEleFengniaoTradeUpdateAPIRequest() *AlibabaEleFengniaoTradeUpdateAPIRequest {
	return poolAlibabaEleFengniaoTradeUpdateAPIRequest.Get().(*AlibabaEleFengniaoTradeUpdateAPIRequest)
}

// ReleaseAlibabaEleFengniaoTradeUpdateAPIRequest 将 AlibabaEleFengniaoTradeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoTradeUpdateAPIRequest(v *AlibabaEleFengniaoTradeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoTradeUpdateAPIRequest.Put(v)
}
