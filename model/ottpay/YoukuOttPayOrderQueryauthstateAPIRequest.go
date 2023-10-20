package ottpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryauthstateAPIRequest 查询连包签约状态 API请求
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
type YoukuOttPayOrderQueryauthstateAPIRequest struct {
	model.Params
	// 原始签约订单号
	_originalCpOrderNo string
}

// NewYoukuOttPayOrderQueryauthstateRequest 初始化YoukuOttPayOrderQueryauthstateAPIRequest对象
func NewYoukuOttPayOrderQueryauthstateRequest() *YoukuOttPayOrderQueryauthstateAPIRequest {
	return &YoukuOttPayOrderQueryauthstateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttPayOrderQueryauthstateAPIRequest) Reset() {
	r._originalCpOrderNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryauthstate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginalCpOrderNo is OriginalCpOrderNo Setter
// 原始签约订单号
func (r *YoukuOttPayOrderQueryauthstateAPIRequest) SetOriginalCpOrderNo(_originalCpOrderNo string) error {
	r._originalCpOrderNo = _originalCpOrderNo
	r.Set("original_cp_order_no", _originalCpOrderNo)
	return nil
}

// GetOriginalCpOrderNo OriginalCpOrderNo Getter
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetOriginalCpOrderNo() string {
	return r._originalCpOrderNo
}

var poolYoukuOttPayOrderQueryauthstateAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttPayOrderQueryauthstateRequest()
	},
}

// GetYoukuOttPayOrderQueryauthstateRequest 从 sync.Pool 获取 YoukuOttPayOrderQueryauthstateAPIRequest
func GetYoukuOttPayOrderQueryauthstateAPIRequest() *YoukuOttPayOrderQueryauthstateAPIRequest {
	return poolYoukuOttPayOrderQueryauthstateAPIRequest.Get().(*YoukuOttPayOrderQueryauthstateAPIRequest)
}

// ReleaseYoukuOttPayOrderQueryauthstateAPIRequest 将 YoukuOttPayOrderQueryauthstateAPIRequest 放入 sync.Pool
func ReleaseYoukuOttPayOrderQueryauthstateAPIRequest(v *YoukuOttPayOrderQueryauthstateAPIRequest) {
	v.Reset()
	poolYoukuOttPayOrderQueryauthstateAPIRequest.Put(v)
}
