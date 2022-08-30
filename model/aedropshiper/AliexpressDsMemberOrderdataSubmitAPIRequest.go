package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsMemberOrderdataSubmitAPIRequest dropshipper数据回流 API请求
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
type AliexpressDsMemberOrderdataSubmitAPIRequest struct {
	model.Params
	// AE product ID
	_aeProductId string
	// SKU sales amount outside the station, to 2 decimal places
	_productAmount string
	// Order sales amount outside the station, keep 2 decimal places
	_orderAmount string
	// Off-site payment time, GMT time, format YYYYMMDD:HHMMSS
	_paytime string
	// AE product SKU information, SKU key-value pair: "200000182:193;200007763:201336100"
	_aeSkuInfo string
	// Commodity site url
	_productUrl string
	// AE order id
	_aeOrderid string
}

// NewAliexpressDsMemberOrderdataSubmitRequest 初始化AliexpressDsMemberOrderdataSubmitAPIRequest对象
func NewAliexpressDsMemberOrderdataSubmitRequest() *AliexpressDsMemberOrderdataSubmitAPIRequest {
	return &AliexpressDsMemberOrderdataSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.member.orderdata.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAeProductId is AeProductId Setter
// AE product ID
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetAeProductId(_aeProductId string) error {
	r._aeProductId = _aeProductId
	r.Set("ae_product_id", _aeProductId)
	return nil
}

// GetAeProductId AeProductId Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetAeProductId() string {
	return r._aeProductId
}

// SetProductAmount is ProductAmount Setter
// SKU sales amount outside the station, to 2 decimal places
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetProductAmount(_productAmount string) error {
	r._productAmount = _productAmount
	r.Set("product_amount", _productAmount)
	return nil
}

// GetProductAmount ProductAmount Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetProductAmount() string {
	return r._productAmount
}

// SetOrderAmount is OrderAmount Setter
// Order sales amount outside the station, keep 2 decimal places
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetOrderAmount(_orderAmount string) error {
	r._orderAmount = _orderAmount
	r.Set("order_amount", _orderAmount)
	return nil
}

// GetOrderAmount OrderAmount Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetOrderAmount() string {
	return r._orderAmount
}

// SetPaytime is Paytime Setter
// Off-site payment time, GMT time, format YYYYMMDD:HHMMSS
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetPaytime(_paytime string) error {
	r._paytime = _paytime
	r.Set("paytime", _paytime)
	return nil
}

// GetPaytime Paytime Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetPaytime() string {
	return r._paytime
}

// SetAeSkuInfo is AeSkuInfo Setter
// AE product SKU information, SKU key-value pair: "200000182:193;200007763:201336100"
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetAeSkuInfo(_aeSkuInfo string) error {
	r._aeSkuInfo = _aeSkuInfo
	r.Set("ae_sku_info", _aeSkuInfo)
	return nil
}

// GetAeSkuInfo AeSkuInfo Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetAeSkuInfo() string {
	return r._aeSkuInfo
}

// SetProductUrl is ProductUrl Setter
// Commodity site url
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetProductUrl(_productUrl string) error {
	r._productUrl = _productUrl
	r.Set("product_url", _productUrl)
	return nil
}

// GetProductUrl ProductUrl Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetProductUrl() string {
	return r._productUrl
}

// SetAeOrderid is AeOrderid Setter
// AE order id
func (r *AliexpressDsMemberOrderdataSubmitAPIRequest) SetAeOrderid(_aeOrderid string) error {
	r._aeOrderid = _aeOrderid
	r.Set("ae_orderid", _aeOrderid)
	return nil
}

// GetAeOrderid AeOrderid Getter
func (r AliexpressDsMemberOrderdataSubmitAPIRequest) GetAeOrderid() string {
	return r._aeOrderid
}
