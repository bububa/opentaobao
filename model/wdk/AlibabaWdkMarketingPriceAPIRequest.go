package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingPriceAPIRequest 促销价签服务 API请求
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
type AlibabaWdkMarketingPriceAPIRequest struct {
	model.Params
	// 单页大小
	_pageSize int64
	// 页码
	_pageIndex int64
	// 商品sku
	_skuCodes []string
	// 门店标识数组
	_shopIds []int64
	// 查询结束时间(sku_codes非空无效)
	_endTime string
	// 查询开始时间(sku_codes非空无效)
	_beginTime string
}

// NewAlibabaWdkMarketingPriceRequest 初始化AlibabaWdkMarketingPriceAPIRequest对象
func NewAlibabaWdkMarketingPriceRequest() *AlibabaWdkMarketingPriceAPIRequest {
	return &AlibabaWdkMarketingPriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingPriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPageSize is PageSize Setter
// 单页大小
func (r *AlibabaWdkMarketingPriceAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *AlibabaWdkMarketingPriceAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetSkuCodes is SkuCodes Setter
// 商品sku
func (r *AlibabaWdkMarketingPriceAPIRequest) SetSkuCodes(_skuCodes []string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetSkuCodes() []string {
	return r._skuCodes
}

// SetShopIds is ShopIds Setter
// 门店标识数组
func (r *AlibabaWdkMarketingPriceAPIRequest) SetShopIds(_shopIds []int64) error {
	r._shopIds = _shopIds
	r.Set("shop_ids", _shopIds)
	return nil
}

// GetShopIds ShopIds Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetShopIds() []int64 {
	return r._shopIds
}

// SetEndTime is EndTime Setter
// 查询结束时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetBeginTime is BeginTime Setter
// 查询开始时间(sku_codes非空无效)
func (r *AlibabaWdkMarketingPriceAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetBeginTime() string {
	return r._beginTime
}
