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
type TmallChannelProductsQueryAPIRequest struct {
    model.Params
    // 商家产品编码
    _productNumber   string
    // 商家SKU编码
    _skuNumber   string
    // 产品Id
    _productIds   []int64
    // 分页大小
    _pageSize   int64
    // 产品线Id
    _productLineId   int64
    // 页码数，从1开始
    _pageNum   int64
}

// 初始化TmallChannelProductsQueryAPIRequest对象
func NewTmallChannelProductsQueryRequest() *TmallChannelProductsQueryAPIRequest{
    return &TmallChannelProductsQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelProductsQueryAPIRequest) GetApiMethodName() string {
    return "tmall.channel.products.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelProductsQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductNumber Setter
// 商家产品编码
func (r *TmallChannelProductsQueryAPIRequest) SetProductNumber(_productNumber string) error {
    r._productNumber = _productNumber
    r.Set("product_number", _productNumber)
    return nil
}

// ProductNumber Getter
func (r TmallChannelProductsQueryAPIRequest) GetProductNumber() string {
    return r._productNumber
}
// SkuNumber Setter
// 商家SKU编码
func (r *TmallChannelProductsQueryAPIRequest) SetSkuNumber(_skuNumber string) error {
    r._skuNumber = _skuNumber
    r.Set("sku_number", _skuNumber)
    return nil
}

// SkuNumber Getter
func (r TmallChannelProductsQueryAPIRequest) GetSkuNumber() string {
    return r._skuNumber
}
// ProductIds Setter
// 产品Id
func (r *TmallChannelProductsQueryAPIRequest) SetProductIds(_productIds []int64) error {
    r._productIds = _productIds
    r.Set("product_ids", _productIds)
    return nil
}

// ProductIds Getter
func (r TmallChannelProductsQueryAPIRequest) GetProductIds() []int64 {
    return r._productIds
}
// PageSize Setter
// 分页大小
func (r *TmallChannelProductsQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TmallChannelProductsQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// ProductLineId Setter
// 产品线Id
func (r *TmallChannelProductsQueryAPIRequest) SetProductLineId(_productLineId int64) error {
    r._productLineId = _productLineId
    r.Set("product_line_id", _productLineId)
    return nil
}

// ProductLineId Getter
func (r TmallChannelProductsQueryAPIRequest) GetProductLineId() int64 {
    return r._productLineId
}
// PageNum Setter
// 页码数，从1开始
func (r *TmallChannelProductsQueryAPIRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TmallChannelProductsQueryAPIRequest) GetPageNum() int64 {
    return r._pageNum
}
