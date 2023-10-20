package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolstairadditemAPIRequest 商品池阶梯商品添加 API请求
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
type AlibabahmmarketingitempoolstairadditemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabahmmarketingitempoolstairadditemRequest 初始化AlibabahmmarketingitempoolstairadditemAPIRequest对象
func NewAlibabahmmarketingitempoolstairadditemRequest() *AlibabahmmarketingitempoolstairadditemAPIRequest {
	return &AlibabahmmarketingitempoolstairadditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolstairadditemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.stair.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolstairadditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolstairadditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabahmmarketingitempoolstairadditemAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitempoolstairadditemAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabahmmarketingitempoolstairadditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingitempoolstairadditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
