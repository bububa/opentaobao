package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategoryaddAPIRequest 商家类目新增接口 API请求
// alibaba.wdk.sku.category.add
//
// 商家类目新增接口
type AlibabawdkskucategoryaddAPIRequest struct {
	model.Params
	// 类目新增请求模型
	_param *CategoryDo
}

// NewAlibabawdkskucategoryaddRequest 初始化AlibabawdkskucategoryaddAPIRequest对象
func NewAlibabawdkskucategoryaddRequest() *AlibabawdkskucategoryaddAPIRequest {
	return &AlibabawdkskucategoryaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucategoryaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucategoryaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucategoryaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 类目新增请求模型
func (r *AlibabawdkskucategoryaddAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskucategoryaddAPIRequest) GetParam() *CategoryDo {
	return r._param
}
