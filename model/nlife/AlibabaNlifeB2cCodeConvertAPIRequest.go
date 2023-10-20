package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// Alibabanlifeb2ccodeconvertAPIRequest b2c转码 API请求
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
type Alibabanlifeb2ccodeconvertAPIRequest struct {
	model.Params
	// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
	_storeId string
	// 商品URL
	_url string
}

// NewAlibabanlifeb2ccodeconvertRequest 初始化Alibabanlifeb2ccodeconvertAPIRequest对象
func NewAlibabanlifeb2ccodeconvertRequest() *Alibabanlifeb2ccodeconvertAPIRequest {
	return &Alibabanlifeb2ccodeconvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r Alibabanlifeb2ccodeconvertAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.code.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r Alibabanlifeb2ccodeconvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r Alibabanlifeb2ccodeconvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
func (r *Alibabanlifeb2ccodeconvertAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r Alibabanlifeb2ccodeconvertAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetUrl is Url Setter
// 商品URL
func (r *Alibabanlifeb2ccodeconvertAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r Alibabanlifeb2ccodeconvertAPIRequest) GetUrl() string {
	return r._url
}
