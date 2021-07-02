package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemPriceUpdateAPIRequest 五道口商品中心价格修改 API请求
// alibaba.wdk.item.price.update
//
// 修改门店商品售价和会员价
type AlibabaWdkItemPriceUpdateAPIRequest struct {
	model.Params
	// 盒马门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 价格，单位是分
	_skuPrice int64
	// 会员价格，单位是分，且不能大于价格
	_skuMemberPrice int64
}

// NewAlibabaWdkItemPriceUpdateRequest 初始化AlibabaWdkItemPriceUpdateAPIRequest对象
func NewAlibabaWdkItemPriceUpdateRequest() *AlibabaWdkItemPriceUpdateAPIRequest {
	return &AlibabaWdkItemPriceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStoreId is StoreId Setter
// 盒马门店id
func (r *AlibabaWdkItemPriceUpdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabaWdkItemPriceUpdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetSkuPrice is SkuPrice Setter
// 价格，单位是分
func (r *AlibabaWdkItemPriceUpdateAPIRequest) SetSkuPrice(_skuPrice int64) error {
	r._skuPrice = _skuPrice
	r.Set("sku_price", _skuPrice)
	return nil
}

// GetSkuPrice SkuPrice Getter
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetSkuPrice() int64 {
	return r._skuPrice
}

// SetSkuMemberPrice is SkuMemberPrice Setter
// 会员价格，单位是分，且不能大于价格
func (r *AlibabaWdkItemPriceUpdateAPIRequest) SetSkuMemberPrice(_skuMemberPrice int64) error {
	r._skuMemberPrice = _skuMemberPrice
	r.Set("sku_member_price", _skuMemberPrice)
	return nil
}

// GetSkuMemberPrice SkuMemberPrice Getter
func (r AlibabaWdkItemPriceUpdateAPIRequest) GetSkuMemberPrice() int64 {
	return r._skuMemberPrice
}
