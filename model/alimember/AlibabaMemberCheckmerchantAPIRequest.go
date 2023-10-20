package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberCheckmerchantAPIRequest 校验商家身份 API请求
// alibaba.member.checkmerchant
//
// 校验商家身份
type AlibabaMemberCheckmerchantAPIRequest struct {
	model.Params
	// 混淆后的商家id
	_openMerchantId string
}

// NewAlibabaMemberCheckmerchantRequest 初始化AlibabaMemberCheckmerchantAPIRequest对象
func NewAlibabaMemberCheckmerchantRequest() *AlibabaMemberCheckmerchantAPIRequest {
	return &AlibabaMemberCheckmerchantAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberCheckmerchantAPIRequest) Reset() {
	r._openMerchantId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberCheckmerchantAPIRequest) GetApiMethodName() string {
	return "alibaba.member.checkmerchant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberCheckmerchantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberCheckmerchantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenMerchantId is OpenMerchantId Setter
// 混淆后的商家id
func (r *AlibabaMemberCheckmerchantAPIRequest) SetOpenMerchantId(_openMerchantId string) error {
	r._openMerchantId = _openMerchantId
	r.Set("open_merchant_id", _openMerchantId)
	return nil
}

// GetOpenMerchantId OpenMerchantId Getter
func (r AlibabaMemberCheckmerchantAPIRequest) GetOpenMerchantId() string {
	return r._openMerchantId
}

var poolAlibabaMemberCheckmerchantAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberCheckmerchantRequest()
	},
}

// GetAlibabaMemberCheckmerchantRequest 从 sync.Pool 获取 AlibabaMemberCheckmerchantAPIRequest
func GetAlibabaMemberCheckmerchantAPIRequest() *AlibabaMemberCheckmerchantAPIRequest {
	return poolAlibabaMemberCheckmerchantAPIRequest.Get().(*AlibabaMemberCheckmerchantAPIRequest)
}

// ReleaseAlibabaMemberCheckmerchantAPIRequest 将 AlibabaMemberCheckmerchantAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberCheckmerchantAPIRequest(v *AlibabaMemberCheckmerchantAPIRequest) {
	v.Reset()
	poolAlibabaMemberCheckmerchantAPIRequest.Put(v)
}
