package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuCategoryAddAPIRequest
商家类目新增接口 API请求
alibaba.wdk.sku.category.add

商家类目新增接口 */
type AlibabaWdkSkuCategoryAddAPIRequest struct {
	model.Params
	// 类目新增请求模型
	_param *CategoryDo
}

// NewAlibabaWdkSkuCategoryAddRequest 初始化AlibabaWdkSkuCategoryAddAPIRequest对象
func NewAlibabaWdkSkuCategoryAddRequest() *AlibabaWdkSkuCategoryAddAPIRequest {
	return &AlibabaWdkSkuCategoryAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCategoryAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCategoryAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 类目新增请求模型
func (r *AlibabaWdkSkuCategoryAddAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaWdkSkuCategoryAddAPIRequest) GetParam() *CategoryDo {
	return r._param
}
