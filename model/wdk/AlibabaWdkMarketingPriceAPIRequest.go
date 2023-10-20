package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingpriceAPIRequest 促销价签服务 API请求
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
type AlibabawdkmarketingpriceAPIRequest struct {
	model.Params
	// 商品sku
	_skuCodes []string
	// 门店标识数组
	_shopIds []string
	// 查询结束时间(sku_codes非空无效)
	_endTime string
	// 查询开始时间(sku_codes非空无效)
	_beginTime string
	// 单页大小
	_pageSize int64
	// 页码
	_pageIndex int64
}

// NewAlibabawdkmarketingpriceRequest 初始化AlibabawdkmarketingpriceAPIRequest对象
func NewAlibabawdkmarketingpriceRequest() *AlibabawdkmarketingpriceAPIRequest {
	return &AlibabawdkmarketingpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCodes is SkuCodes Setter
// 商品sku
func (r *AlibabawdkmarketingpriceAPIRequest) SetSkuCodes(_skuCodes []string) error {
	r._skuCodes = _skuCodes
	r.Set("sku_codes", _skuCodes)
	return nil
}

// GetSkuCodes SkuCodes Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetSkuCodes() []string {
	return r._skuCodes
}

// SetShopIds is ShopIds Setter
// 门店标识数组
func (r *AlibabawdkmarketingpriceAPIRequest) SetShopIds(_shopIds []string) error {
	r._shopIds = _shopIds
	r.Set("shop_ids", _shopIds)
	return nil
}

// GetShopIds ShopIds Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetShopIds() []string {
	return r._shopIds
}

// SetEndTime is EndTime Setter
// 查询结束时间(sku_codes非空无效)
func (r *AlibabawdkmarketingpriceAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetBeginTime is BeginTime Setter
// 查询开始时间(sku_codes非空无效)
func (r *AlibabawdkmarketingpriceAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetPageSize is PageSize Setter
// 单页大小
func (r *AlibabawdkmarketingpriceAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *AlibabawdkmarketingpriceAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlibabawdkmarketingpriceAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
