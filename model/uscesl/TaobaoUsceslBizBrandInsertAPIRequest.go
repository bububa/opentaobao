package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizbrandinsertAPIRequest 新增电子价签商家 API请求
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
type TaobaousceslbizbrandinsertAPIRequest struct {
	model.Params
	// 商家名称
	_brandName string
	// 商家外部编号
	_brandOutCode string
}

// NewTaobaousceslbizbrandinsertRequest 初始化TaobaousceslbizbrandinsertAPIRequest对象
func NewTaobaousceslbizbrandinsertRequest() *TaobaousceslbizbrandinsertAPIRequest {
	return &TaobaousceslbizbrandinsertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizbrandinsertAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.brand.insert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizbrandinsertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizbrandinsertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandName is BrandName Setter
// 商家名称
func (r *TaobaousceslbizbrandinsertAPIRequest) SetBrandName(_brandName string) error {
	r._brandName = _brandName
	r.Set("brand_name", _brandName)
	return nil
}

// GetBrandName BrandName Getter
func (r TaobaousceslbizbrandinsertAPIRequest) GetBrandName() string {
	return r._brandName
}

// SetBrandOutCode is BrandOutCode Setter
// 商家外部编号
func (r *TaobaousceslbizbrandinsertAPIRequest) SetBrandOutCode(_brandOutCode string) error {
	r._brandOutCode = _brandOutCode
	r.Set("brand_out_code", _brandOutCode)
	return nil
}

// GetBrandOutCode BrandOutCode Getter
func (r TaobaousceslbizbrandinsertAPIRequest) GetBrandOutCode() string {
	return r._brandOutCode
}
