package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道中心-查询产品列表 API请求
tmall.channel.products.query

渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
*/
type TmallChannelProductsQueryRequest struct {
    model.Params
    // 商家产品编码
    productNumber   string
    // 商家SKU编码
    skuNumber   string
    // 产品Id
    productIds   []int64
    // 分页大小
    pageSize   int64
    // 产品线Id
    productLineId   int64
    // 页码数，从1开始
    pageNum   int64
}

// 初始化TmallChannelProductsQueryRequest对象
func NewTmallChannelProductsQueryRequest() *TmallChannelProductsQueryRequest{
    return &TmallChannelProductsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelProductsQueryRequest) GetApiMethodName() string {
    return "tmall.channel.products.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelProductsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductNumber Setter
// 商家产品编码
func (r *TmallChannelProductsQueryRequest) SetProductNumber(productNumber string) error {
    r.productNumber = productNumber
    r.Set("product_number", productNumber)
    return nil
}

// ProductNumber Getter
func (r TmallChannelProductsQueryRequest) GetProductNumber() string {
    return r.productNumber
}
// SkuNumber Setter
// 商家SKU编码
func (r *TmallChannelProductsQueryRequest) SetSkuNumber(skuNumber string) error {
    r.skuNumber = skuNumber
    r.Set("sku_number", skuNumber)
    return nil
}

// SkuNumber Getter
func (r TmallChannelProductsQueryRequest) GetSkuNumber() string {
    return r.skuNumber
}
// ProductIds Setter
// 产品Id
func (r *TmallChannelProductsQueryRequest) SetProductIds(productIds []int64) error {
    r.productIds = productIds
    r.Set("product_ids", productIds)
    return nil
}

// ProductIds Getter
func (r TmallChannelProductsQueryRequest) GetProductIds() []int64 {
    return r.productIds
}
// PageSize Setter
// 分页大小
func (r *TmallChannelProductsQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelProductsQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// ProductLineId Setter
// 产品线Id
func (r *TmallChannelProductsQueryRequest) SetProductLineId(productLineId int64) error {
    r.productLineId = productLineId
    r.Set("product_line_id", productLineId)
    return nil
}

// ProductLineId Getter
func (r TmallChannelProductsQueryRequest) GetProductLineId() int64 {
    return r.productLineId
}
// PageNum Setter
// 页码数，从1开始
func (r *TmallChannelProductsQueryRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TmallChannelProductsQueryRequest) GetPageNum() int64 {
    return r.pageNum
}
