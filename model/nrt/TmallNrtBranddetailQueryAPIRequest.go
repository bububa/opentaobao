package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtbranddetailqueryAPIRequest 品牌详情查询 API请求
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
type TmallnrtbranddetailqueryAPIRequest struct {
	model.Params
	// 品牌id
	_brandId int64
}

// NewTmallnrtbranddetailqueryRequest 初始化TmallnrtbranddetailqueryAPIRequest对象
func NewTmallnrtbranddetailqueryRequest() *TmallnrtbranddetailqueryAPIRequest {
	return &TmallnrtbranddetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtbranddetailqueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.branddetail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtbranddetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtbranddetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandId is BrandId Setter
// 品牌id
func (r *TmallnrtbranddetailqueryAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallnrtbranddetailqueryAPIRequest) GetBrandId() int64 {
	return r._brandId
}
