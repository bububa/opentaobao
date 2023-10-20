package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuupdateAPIRequest 更新渠道商品 API请求
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
type AlibabawdkskuchannelskuupdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []ChannelSkuDo
}

// NewAlibabawdkskuchannelskuupdateRequest 初始化AlibabawdkskuchannelskuupdateAPIRequest对象
func NewAlibabawdkskuchannelskuupdateRequest() *AlibabawdkskuchannelskuupdateAPIRequest {
	return &AlibabawdkskuchannelskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuchannelskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuchannelskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuchannelskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabawdkskuchannelskuupdateAPIRequest) SetParamList(_paramList []ChannelSkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabawdkskuchannelskuupdateAPIRequest) GetParamList() []ChannelSkuDo {
	return r._paramList
}
