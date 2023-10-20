package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingPriceAPIRequest 促销价签服务 API请求
// alibaba.wdk.marketing.price
//
// 获取营销-促销商品中的实时价格
type AlibabaWdkMarketingPriceAPIRequest struct {
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

// NewAlibabaWdkMarketingPriceRequest 初始化AlibabaWdkMarketingPriceAPIRequest对象
func NewAlibabaWdkMarketingPriceRequest() *AlibabaWdkMarketingPriceAPIRequest {
	return &AlibabaWdkMarketingPriceAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingPriceAPIRequest) Reset() {
	r._skuCodes = r._skuCodes[:0]
	r._shopIds = r._shopIds[:0]
	r._endTime = ""
	r._beginTime = ""
	r._pageSize = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingPriceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingPriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingPriceAPIRequest) GetRawParams() model.Params {
	return r.Params
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
func (r *AlibabaWdkMarketingPriceAPIRequest) SetShopIds(_shopIds []string) error {
	r._shopIds = _shopIds
	r.Set("shop_ids", _shopIds)
	return nil
}

// GetShopIds ShopIds Getter
func (r AlibabaWdkMarketingPriceAPIRequest) GetShopIds() []string {
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

var poolAlibabaWdkMarketingPriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingPriceRequest()
	},
}

// GetAlibabaWdkMarketingPriceRequest 从 sync.Pool 获取 AlibabaWdkMarketingPriceAPIRequest
func GetAlibabaWdkMarketingPriceAPIRequest() *AlibabaWdkMarketingPriceAPIRequest {
	return poolAlibabaWdkMarketingPriceAPIRequest.Get().(*AlibabaWdkMarketingPriceAPIRequest)
}

// ReleaseAlibabaWdkMarketingPriceAPIRequest 将 AlibabaWdkMarketingPriceAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingPriceAPIRequest(v *AlibabaWdkMarketingPriceAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingPriceAPIRequest.Put(v)
}
