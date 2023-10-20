package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategorydeleteAPIRequest 商家类目删除接口 API请求
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
type AlibabawdkskucategorydeleteAPIRequest struct {
	model.Params
	// 类目删除请求模型
	_param *CategoryDo
}

// NewAlibabawdkskucategorydeleteRequest 初始化AlibabawdkskucategorydeleteAPIRequest对象
func NewAlibabawdkskucategorydeleteRequest() *AlibabawdkskucategorydeleteAPIRequest {
	return &AlibabawdkskucategorydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucategorydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucategorydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucategorydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 类目删除请求模型
func (r *AlibabawdkskucategorydeleteAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskucategorydeleteAPIRequest) GetParam() *CategoryDo {
	return r._param
}
