package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategoryupdateAPIRequest 商家类目修改接口 API请求
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
type AlibabawdkskucategoryupdateAPIRequest struct {
	model.Params
	// 更新请求模型
	_param *CategoryDo
}

// NewAlibabawdkskucategoryupdateRequest 初始化AlibabawdkskucategoryupdateAPIRequest对象
func NewAlibabawdkskucategoryupdateRequest() *AlibabawdkskucategoryupdateAPIRequest {
	return &AlibabawdkskucategoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucategoryupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucategoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucategoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新请求模型
func (r *AlibabawdkskucategoryupdateAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskucategoryupdateAPIRequest) GetParam() *CategoryDo {
	return r._param
}
