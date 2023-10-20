package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucombineskuupdateAPIRequest 组合商品更新接口 API请求
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
type AlibabawdkskucombineskuupdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// NewAlibabawdkskucombineskuupdateRequest 初始化AlibabawdkskucombineskuupdateAPIRequest对象
func NewAlibabawdkskucombineskuupdateRequest() *AlibabawdkskucombineskuupdateAPIRequest {
	return &AlibabawdkskucombineskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucombineskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucombineskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucombineskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabawdkskucombineskuupdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabawdkskucombineskuupdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
