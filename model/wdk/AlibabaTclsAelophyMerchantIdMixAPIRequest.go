package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantIdMixAPIRequest 商家用户id混淆 API请求
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
type AlibabaTclsAelophyMerchantIdMixAPIRequest struct {
	model.Params
	// 商家用户id
	_unionUid string
}

// NewAlibabaTclsAelophyMerchantIdMixRequest 初始化AlibabaTclsAelophyMerchantIdMixAPIRequest对象
func NewAlibabaTclsAelophyMerchantIdMixRequest() *AlibabaTclsAelophyMerchantIdMixAPIRequest {
	return &AlibabaTclsAelophyMerchantIdMixAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantIdMixAPIRequest) Reset() {
	r._unionUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.id.mix"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnionUid is UnionUid Setter
// 商家用户id
func (r *AlibabaTclsAelophyMerchantIdMixAPIRequest) SetUnionUid(_unionUid string) error {
	r._unionUid = _unionUid
	r.Set("union_uid", _unionUid)
	return nil
}

// GetUnionUid UnionUid Getter
func (r AlibabaTclsAelophyMerchantIdMixAPIRequest) GetUnionUid() string {
	return r._unionUid
}

var poolAlibabaTclsAelophyMerchantIdMixAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantIdMixRequest()
	},
}

// GetAlibabaTclsAelophyMerchantIdMixRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantIdMixAPIRequest
func GetAlibabaTclsAelophyMerchantIdMixAPIRequest() *AlibabaTclsAelophyMerchantIdMixAPIRequest {
	return poolAlibabaTclsAelophyMerchantIdMixAPIRequest.Get().(*AlibabaTclsAelophyMerchantIdMixAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantIdMixAPIRequest 将 AlibabaTclsAelophyMerchantIdMixAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantIdMixAPIRequest(v *AlibabaTclsAelophyMerchantIdMixAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantIdMixAPIRequest.Put(v)
}
