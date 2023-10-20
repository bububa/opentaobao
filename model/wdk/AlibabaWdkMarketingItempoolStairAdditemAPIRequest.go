package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolstairadditemAPIRequest 商品池阶梯商品添加 API请求
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabawdkmarketingitempoolstairadditemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingitempoolstairadditemRequest 初始化AlibabawdkmarketingitempoolstairadditemAPIRequest对象
func NewAlibabawdkmarketingitempoolstairadditemRequest() *AlibabawdkmarketingitempoolstairadditemAPIRequest {
	return &AlibabawdkmarketingitempoolstairadditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolstairadditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.stair.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolstairadditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolstairadditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabawdkmarketingitempoolstairadditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitempoolstairadditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabawdkmarketingitempoolstairadditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingitempoolstairadditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
