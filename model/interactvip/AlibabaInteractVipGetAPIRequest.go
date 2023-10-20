package interactvip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractVipGetAPIRequest 会员淘气值获取 API请求
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
type AlibabaInteractVipGetAPIRequest struct {
	model.Params
}

// NewAlibabaInteractVipGetRequest 初始化AlibabaInteractVipGetAPIRequest对象
func NewAlibabaInteractVipGetRequest() *AlibabaInteractVipGetAPIRequest {
	return &AlibabaInteractVipGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractVipGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractVipGetAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.vip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractVipGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractVipGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractVipGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractVipGetRequest()
	},
}

// GetAlibabaInteractVipGetRequest 从 sync.Pool 获取 AlibabaInteractVipGetAPIRequest
func GetAlibabaInteractVipGetAPIRequest() *AlibabaInteractVipGetAPIRequest {
	return poolAlibabaInteractVipGetAPIRequest.Get().(*AlibabaInteractVipGetAPIRequest)
}

// ReleaseAlibabaInteractVipGetAPIRequest 将 AlibabaInteractVipGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractVipGetAPIRequest(v *AlibabaInteractVipGetAPIRequest) {
	v.Reset()
	poolAlibabaInteractVipGetAPIRequest.Put(v)
}
