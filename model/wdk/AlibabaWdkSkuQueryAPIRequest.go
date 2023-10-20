package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuqueryAPIRequest 查询商品 API请求
// alibaba.wdk.sku.query
//
// 查询商品
type AlibabawdkskuqueryAPIRequest struct {
	model.Params
	// 入参
	_param *SkuQueryDo
}

// NewAlibabawdkskuqueryRequest 初始化AlibabawdkskuqueryAPIRequest对象
func NewAlibabawdkskuqueryRequest() *AlibabawdkskuqueryAPIRequest {
	return &AlibabawdkskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabawdkskuqueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskuqueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}
