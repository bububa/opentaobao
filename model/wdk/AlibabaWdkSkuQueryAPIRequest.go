package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuQueryAPIRequest 查询商品 API请求
// alibaba.wdk.sku.query
//
// 查询商品
type AlibabaWdkSkuQueryAPIRequest struct {
	model.Params
	// 入参
	_param *SkuQueryDo
}

// NewAlibabaWdkSkuQueryRequest 初始化AlibabaWdkSkuQueryAPIRequest对象
func NewAlibabaWdkSkuQueryRequest() *AlibabaWdkSkuQueryAPIRequest {
	return &AlibabaWdkSkuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaWdkSkuQueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuQueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}
