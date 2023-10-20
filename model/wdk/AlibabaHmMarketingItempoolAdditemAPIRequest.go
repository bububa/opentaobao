package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempooladditemAPIRequest 增加商品池里面的商品 API请求
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabahmmarketingitempooladditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabahmmarketingitempooladditemRequest 初始化AlibabahmmarketingitempooladditemAPIRequest对象
func NewAlibabahmmarketingitempooladditemRequest() *AlibabahmmarketingitempooladditemAPIRequest {
	return &AlibabahmmarketingitempooladditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempooladditemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempooladditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempooladditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabahmmarketingitempooladditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitempooladditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabahmmarketingitempooladditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingitempooladditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
