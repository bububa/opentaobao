package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuQueryAPIRequest 组合商品查询接口 API请求
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
type AlibabaWdkSkuCombineskuQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param *SkuQueryDo
}

// NewAlibabaWdkSkuCombineskuQueryRequest 初始化AlibabaWdkSkuCombineskuQueryAPIRequest对象
func NewAlibabaWdkSkuCombineskuQueryRequest() *AlibabaWdkSkuCombineskuQueryAPIRequest {
	return &AlibabaWdkSkuCombineskuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuQueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}
