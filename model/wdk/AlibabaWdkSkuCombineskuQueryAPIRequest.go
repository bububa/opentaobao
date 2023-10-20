package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucombineskuqueryAPIRequest 组合商品查询接口 API请求
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
type AlibabawdkskucombineskuqueryAPIRequest struct {
	model.Params
	// 请求参数
	_param *SkuQueryDo
}

// NewAlibabawdkskucombineskuqueryRequest 初始化AlibabawdkskucombineskuqueryAPIRequest对象
func NewAlibabawdkskucombineskuqueryRequest() *AlibabawdkskucombineskuqueryAPIRequest {
	return &AlibabawdkskucombineskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskucombineskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskucombineskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskucombineskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabawdkskucombineskuqueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskucombineskuqueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}
