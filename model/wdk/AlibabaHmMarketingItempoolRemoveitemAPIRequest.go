package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolremoveitemAPIRequest 移除商品池里面的商品 API请求
// alibaba.hm.marketing.itempool.removeitem
//
// 移除商品池里面的商品
type AlibabahmmarketingitempoolremoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabahmmarketingitempoolremoveitemRequest 初始化AlibabahmmarketingitempoolremoveitemAPIRequest对象
func NewAlibabahmmarketingitempoolremoveitemRequest() *AlibabahmmarketingitempoolremoveitemAPIRequest {
	return &AlibabahmmarketingitempoolremoveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolremoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolremoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolremoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabahmmarketingitempoolremoveitemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitempoolremoveitemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabahmmarketingitempoolremoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingitempoolremoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
