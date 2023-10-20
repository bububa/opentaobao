package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantidmixAPIRequest 商家用户id混淆 API请求
// alibaba.tcls.aelophy.merchant.id.mix
//
// 商家用户id混淆
type AlibabatclsaelophymerchantidmixAPIRequest struct {
	model.Params
	// 商家用户id
	_unionUid string
}

// NewAlibabatclsaelophymerchantidmixRequest 初始化AlibabatclsaelophymerchantidmixAPIRequest对象
func NewAlibabatclsaelophymerchantidmixRequest() *AlibabatclsaelophymerchantidmixAPIRequest {
	return &AlibabatclsaelophymerchantidmixAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantidmixAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.id.mix"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantidmixAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantidmixAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnionUid is UnionUid Setter
// 商家用户id
func (r *AlibabatclsaelophymerchantidmixAPIRequest) SetUnionUid(_unionUid string) error {
	r._unionUid = _unionUid
	r.Set("union_uid", _unionUid)
	return nil
}

// GetUnionUid UnionUid Getter
func (r AlibabatclsaelophymerchantidmixAPIRequest) GetUnionUid() string {
	return r._unionUid
}
