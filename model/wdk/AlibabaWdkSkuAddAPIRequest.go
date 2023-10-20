package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuaddAPIRequest 新增商品 API请求
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
type AlibabawdkskuaddAPIRequest struct {
	model.Params
	// 商品列表
	_paramList []SkuDo
}

// NewAlibabawdkskuaddRequest 初始化AlibabawdkskuaddAPIRequest对象
func NewAlibabawdkskuaddRequest() *AlibabawdkskuaddAPIRequest {
	return &AlibabawdkskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 商品列表
func (r *AlibabawdkskuaddAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabawdkskuaddAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
