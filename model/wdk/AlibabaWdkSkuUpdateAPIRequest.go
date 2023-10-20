package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuupdateAPIRequest 更新商品 API请求
// alibaba.wdk.sku.update
//
// 开放商品更新接口
type AlibabawdkskuupdateAPIRequest struct {
	model.Params
	// 参数列表
	_paramList []SkuDo
}

// NewAlibabawdkskuupdateRequest 初始化AlibabawdkskuupdateAPIRequest对象
func NewAlibabawdkskuupdateRequest() *AlibabawdkskuupdateAPIRequest {
	return &AlibabawdkskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 参数列表
func (r *AlibabawdkskuupdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabawdkskuupdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}
