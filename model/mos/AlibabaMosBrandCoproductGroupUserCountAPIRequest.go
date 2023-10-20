package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosbrandcoproductgroupusercountAPIRequest 按照查询条件统计总数 API请求
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
type AlibabamosbrandcoproductgroupusercountAPIRequest struct {
	model.Params
	// 查询条件
	_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam
}

// NewAlibabamosbrandcoproductgroupusercountRequest 初始化AlibabamosbrandcoproductgroupusercountAPIRequest对象
func NewAlibabamosbrandcoproductgroupusercountRequest() *AlibabamosbrandcoproductgroupusercountAPIRequest {
	return &AlibabamosbrandcoproductgroupusercountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosbrandcoproductgroupusercountAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.brand.coproduct.group.user.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosbrandcoproductgroupusercountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosbrandcoproductgroupusercountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandCoProductGroupUserQueryParam is BrandCoProductGroupUserQueryParam Setter
// 查询条件
func (r *AlibabamosbrandcoproductgroupusercountAPIRequest) SetBrandCoProductGroupUserQueryParam(_brandCoProductGroupUserQueryParam *BrandCoProductGroupUserQueryParam) error {
	r._brandCoProductGroupUserQueryParam = _brandCoProductGroupUserQueryParam
	r.Set("brand_co_product_group_user_query_param", _brandCoProductGroupUserQueryParam)
	return nil
}

// GetBrandCoProductGroupUserQueryParam BrandCoProductGroupUserQueryParam Getter
func (r AlibabamosbrandcoproductgroupusercountAPIRequest) GetBrandCoProductGroupUserQueryParam() *BrandCoProductGroupUserQueryParam {
	return r._brandCoProductGroupUserQueryParam
}
