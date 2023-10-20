package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCurrentpriceQueryAPIRequest 查询商品当前价格 API请求
// alibaba.wdk.item.currentprice.query
//
// 通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量&lt;=20,返回结果key为skuCode
type AlibabaWdkItemCurrentpriceQueryAPIRequest struct {
	model.Params
	// sku编码列表
	_skuCodes []string
	// 渠道
	_orderChannelCode string
	// 渠道店id
	_shopId int64
}

// NewAlibabaWdkItemCurrentpriceQueryRequest 初始化AlibabaWdkItemCurrentpriceQueryAPIRequest对象
func NewAlibabaWdkItemCurrentpriceQueryRequest() *AlibabaWdkItemCurrentpriceQueryAPIRequest {
	return &AlibabaWdkItemCurrentpriceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.currentprice.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodes is SkuCodes Setter
// sku编码列表
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetSkuCodes(_skuCodes []string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetSkuCodes() []string {
	return r._skuCodes
}

// SetOrderChannelCode is OrderChannelCode Setter
// 渠道
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetOrderChannelCode(_orderChannelCode string) error {
	r._orderChannelCode = _orderChannelCode
	r.Set("order_channel_code", _orderChannelCode)
	return nil
}

// GetOrderChannelCode OrderChannelCode Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetOrderChannelCode() string {
	return r._orderChannelCode
}

// SetShopId is ShopId Setter
// 渠道店id
func (r *AlibabaWdkItemCurrentpriceQueryAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkItemCurrentpriceQueryAPIRequest) GetShopId() int64 {
	return r._shopId
}
