package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallchannelproductsqueryAPIRequest 渠道中心-查询产品列表 API请求
// tmall.channel.products.query
//
// 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallchannelproductsqueryAPIRequest struct {
	model.Params
	// 产品Id
	_productIds []int64
	// 商家产品编码
	_productNumber string
	// 商家SKU编码
	_skuNumber string
	// 分页大小
	_pageSize int64
	// 产品线Id
	_productLineId int64
	// 页码数，从1开始
	_pageNum int64
}

// NewTmallchannelproductsqueryRequest 初始化TmallchannelproductsqueryAPIRequest对象
func NewTmallchannelproductsqueryRequest() *TmallchannelproductsqueryAPIRequest {
	return &TmallchannelproductsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallchannelproductsqueryAPIRequest) GetApiMethodName() string {
	return "tmall.channel.products.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallchannelproductsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallchannelproductsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIds is ProductIds Setter
// 产品Id
func (r *TmallchannelproductsqueryAPIRequest) SetProductIds(_productIds []int64) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r TmallchannelproductsqueryAPIRequest) GetProductIds() []int64 {
	return r._productIds
}

// SetProductNumber is ProductNumber Setter
// 商家产品编码
func (r *TmallchannelproductsqueryAPIRequest) SetProductNumber(_productNumber string) error {
	r._productNumber = _productNumber
	r.Set("product_number", _productNumber)
	return nil
}

// GetProductNumber ProductNumber Getter
func (r TmallchannelproductsqueryAPIRequest) GetProductNumber() string {
	return r._productNumber
}

// SetSkuNumber is SkuNumber Setter
// 商家SKU编码
func (r *TmallchannelproductsqueryAPIRequest) SetSkuNumber(_skuNumber string) error {
	r._skuNumber = _skuNumber
	r.Set("sku_number", _skuNumber)
	return nil
}

// GetSkuNumber SkuNumber Getter
func (r TmallchannelproductsqueryAPIRequest) GetSkuNumber() string {
	return r._skuNumber
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallchannelproductsqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallchannelproductsqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetProductLineId is ProductLineId Setter
// 产品线Id
func (r *TmallchannelproductsqueryAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TmallchannelproductsqueryAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}

// SetPageNum is PageNum Setter
// 页码数，从1开始
func (r *TmallchannelproductsqueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallchannelproductsqueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
