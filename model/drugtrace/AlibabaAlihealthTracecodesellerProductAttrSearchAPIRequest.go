package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest 根据商品id获取商品属性 API请求
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId int64
	// 货品id
	_tracUserProductInfoId int64
}

// NewAlibabaAlihealthTracecodesellerProductAttrSearchRequest 初始化AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerProductAttrSearchRequest() *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) Reset() {
	r._entInfoId = 0
	r._tracUserProductInfoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.product.attr.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntInfoId is EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetTracUserProductInfoId is TracUserProductInfoId Setter
// 货品id
func (r *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) SetTracUserProductInfoId(_tracUserProductInfoId int64) error {
	r._tracUserProductInfoId = _tracUserProductInfoId
	r.Set("trac_user_product_info_id", _tracUserProductInfoId)
	return nil
}

// GetTracUserProductInfoId TracUserProductInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) GetTracUserProductInfoId() int64 {
	return r._tracUserProductInfoId
}

var poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerProductAttrSearchRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerProductAttrSearchRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest
func GetAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest() *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest {
	return poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest.Get().(*AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest 将 AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest(v *AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest.Put(v)
}
