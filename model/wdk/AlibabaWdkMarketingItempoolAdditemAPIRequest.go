package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempooladditemAPIRequest 增加商品池里面的商品 API请求
// alibaba.wdk.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabawdkmarketingitempooladditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingitempooladditemRequest 初始化AlibabawdkmarketingitempooladditemAPIRequest对象
func NewAlibabawdkmarketingitempooladditemRequest() *AlibabawdkmarketingitempooladditemAPIRequest {
	return &AlibabawdkmarketingitempooladditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempooladditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempooladditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempooladditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingitempooladditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitempooladditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabawdkmarketingitempooladditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingitempooladditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
