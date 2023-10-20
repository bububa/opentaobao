package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategoryqueryAPIRequest 商家类目获取接口 API请求
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
type AlibabawdkskucategoryqueryAPIRequest struct {
	model.Params
	// 查询类目请模型
	_param *CategoryDo
}

// NewAlibabawdkskucategoryqueryRequest 初始化AlibabawdkskucategoryqueryAPIRequest对象
func NewAlibabawdkskucategoryqueryRequest() *AlibabawdkskucategoryqueryAPIRequest {
	return &AlibabawdkskucategoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucategoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.category.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucategoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucategoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询类目请模型
func (r *AlibabawdkskucategoryqueryAPIRequest) SetParam(_param *CategoryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskucategoryqueryAPIRequest) GetParam() *CategoryDo {
	return r._param
}
