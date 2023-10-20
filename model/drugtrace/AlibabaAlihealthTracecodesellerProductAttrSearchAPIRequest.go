package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerproductattrsearchAPIRequest 根据商品id获取商品属性 API请求
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
type AlibabaalihealthtracecodesellerproductattrsearchAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId int64
	// 货品id
	_tracUserProductInfoId int64
}

// NewAlibabaalihealthtracecodesellerproductattrsearchRequest 初始化AlibabaalihealthtracecodesellerproductattrsearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerproductattrsearchRequest() *AlibabaalihealthtracecodesellerproductattrsearchAPIRequest {
	return &AlibabaalihealthtracecodesellerproductattrsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.product.attr.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntInfoId is EntInfoId Setter
// 企业id
func (r *AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetTracUserProductInfoId is TracUserProductInfoId Setter
// 货品id
func (r *AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) SetTracUserProductInfoId(_tracUserProductInfoId int64) error {
	r._tracUserProductInfoId = _tracUserProductInfoId
	r.Set("trac_user_product_info_id", _tracUserProductInfoId)
	return nil
}

// GetTracUserProductInfoId TracUserProductInfoId Getter
func (r AlibabaalihealthtracecodesellerproductattrsearchAPIRequest) GetTracUserProductInfoId() int64 {
	return r._tracUserProductInfoId
}
