package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucombineskuaddAPIRequest 组合商品新增接口 API请求
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
type AlibabawdkskucombineskuaddAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// NewAlibabawdkskucombineskuaddRequest 初始化AlibabawdkskucombineskuaddAPIRequest对象
func NewAlibabawdkskucombineskuaddRequest() *AlibabawdkskucombineskuaddAPIRequest {
	return &AlibabawdkskucombineskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucombineskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucombineskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucombineskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabawdkskucombineskuaddAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabawdkskucombineskuaddAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
