package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuqueryAPIRequest 查询渠道商品 API请求
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
type AlibabawdkskuchannelskuqueryAPIRequest struct {
	model.Params
	// 查询渠道商品的入参
	_param *ChannelSkuQueryDo
}

// NewAlibabawdkskuchannelskuqueryRequest 初始化AlibabawdkskuchannelskuqueryAPIRequest对象
func NewAlibabawdkskuchannelskuqueryRequest() *AlibabawdkskuchannelskuqueryAPIRequest {
	return &AlibabawdkskuchannelskuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuchannelskuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuchannelskuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuchannelskuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询渠道商品的入参
func (r *AlibabawdkskuchannelskuqueryAPIRequest) SetParam(_param *ChannelSkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkskuchannelskuqueryAPIRequest) GetParam() *ChannelSkuQueryDo {
	return r._param
}
