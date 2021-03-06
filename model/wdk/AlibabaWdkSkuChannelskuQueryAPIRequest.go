package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuQueryAPIRequest 查询渠道商品 API请求
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
type AlibabaWdkSkuChannelskuQueryAPIRequest struct {
	model.Params
	// 查询渠道商品的入参
	_param *ChannelSkuQueryDo
}

// NewAlibabaWdkSkuChannelskuQueryRequest 初始化AlibabaWdkSkuChannelskuQueryAPIRequest对象
func NewAlibabaWdkSkuChannelskuQueryRequest() *AlibabaWdkSkuChannelskuQueryAPIRequest {
	return &AlibabaWdkSkuChannelskuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 查询渠道商品的入参
func (r *AlibabaWdkSkuChannelskuQueryAPIRequest) SetParam(_param *ChannelSkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetParam() *ChannelSkuQueryDo {
	return r._param
}
