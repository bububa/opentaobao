package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolRemoveitemAPIRequest 移除商品池里面的商品 API请求
// alibaba.wdk.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabaWdkMarketingItempoolRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolRemoveitemRequest 初始化AlibabaWdkMarketingItempoolRemoveitemAPIRequest对象
func NewAlibabaWdkMarketingItempoolRemoveitemRequest() *AlibabaWdkMarketingItempoolRemoveitemAPIRequest {
	return &AlibabaWdkMarketingItempoolRemoveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// Set is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
