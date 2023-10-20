package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosbrandcoproductgroupuserqueryAPIRequest 按照条件查询分页数据 API请求
// alibaba.mos.brand.coproduct.group.user.query
//
// 按照条件查询分页数据
type AlibabamosbrandcoproductgroupuserqueryAPIRequest struct {
	model.Params
	// 查询条件
	_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam
}

// NewAlibabamosbrandcoproductgroupuserqueryRequest 初始化AlibabamosbrandcoproductgroupuserqueryAPIRequest对象
func NewAlibabamosbrandcoproductgroupuserqueryRequest() *AlibabamosbrandcoproductgroupuserqueryAPIRequest {
	return &AlibabamosbrandcoproductgroupuserqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosbrandcoproductgroupuserqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.brand.coproduct.group.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosbrandcoproductgroupuserqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosbrandcoproductgroupuserqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandCoProductGroupUserQueryParam is BrandCoProductGroupUserQueryParam Setter
// 查询条件
func (r *AlibabamosbrandcoproductgroupuserqueryAPIRequest) SetBrandCoProductGroupUserQueryParam(_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam) error {
	r._brandCoProductGroupUserQueryParam = _brandCoProductGroupUserQueryParam
	r.Set("brand_co_product_group_user_query_param", _brandCoProductGroupUserQueryParam)
	return nil
}

// GetBrandCoProductGroupUserQueryParam BrandCoProductGroupUserQueryParam Getter
func (r AlibabamosbrandcoproductgroupuserqueryAPIRequest) GetBrandCoProductGroupUserQueryParam() *BrandCoProductGroupUserQueryParam {
	return r._brandCoProductGroupUserQueryParam
}
