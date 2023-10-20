package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocgetproductbyscancodeAPIRequest POS商品查询接口 API请求
// alibaba.mj.oc.getproductbyscancode
//
// 此API用于在银泰商场中，POS端扫码获取商品信息
type AlibabamjocgetproductbyscancodeAPIRequest struct {
	model.Params
	// 码, 对应的信息可能是款号，也有可能是具体的某一个商品
	_code string
	// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
	_codeType string
	// 专柜编码
	_shopCode string
	// 门店编码
	_storeCode string
}

// NewAlibabamjocgetproductbyscancodeRequest 初始化AlibabamjocgetproductbyscancodeAPIRequest对象
func NewAlibabamjocgetproductbyscancodeRequest() *AlibabamjocgetproductbyscancodeAPIRequest {
	return &AlibabamjocgetproductbyscancodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.getproductbyscancode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 码, 对应的信息可能是款号，也有可能是具体的某一个商品
func (r *AlibabamjocgetproductbyscancodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetCode() string {
	return r._code
}

// SetCodeType is CodeType Setter
// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
func (r *AlibabamjocgetproductbyscancodeAPIRequest) SetCodeType(_codeType string) error {
	r._codeType = _codeType
	r.Set("code_type", _codeType)
	return nil
}

// GetCodeType CodeType Getter
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetCodeType() string {
	return r._codeType
}

// SetShopCode is ShopCode Setter
// 专柜编码
func (r *AlibabamjocgetproductbyscancodeAPIRequest) SetShopCode(_shopCode string) error {
	r._shopCode = _shopCode
	r.Set("shop_code", _shopCode)
	return nil
}

// GetShopCode ShopCode Getter
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetShopCode() string {
	return r._shopCode
}

// SetStoreCode is StoreCode Setter
// 门店编码
func (r *AlibabamjocgetproductbyscancodeAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabamjocgetproductbyscancodeAPIRequest) GetStoreCode() string {
	return r._storeCode
}
