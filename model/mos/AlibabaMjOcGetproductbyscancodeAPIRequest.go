package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcGetproductbyscancodeAPIRequest POS商品查询接口 API请求
// alibaba.mj.oc.getproductbyscancode
//
// 此API用于在银泰商场中，POS端扫码获取商品信息
type AlibabaMjOcGetproductbyscancodeAPIRequest struct {
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

// NewAlibabaMjOcGetproductbyscancodeRequest 初始化AlibabaMjOcGetproductbyscancodeAPIRequest对象
func NewAlibabaMjOcGetproductbyscancodeRequest() *AlibabaMjOcGetproductbyscancodeAPIRequest {
	return &AlibabaMjOcGetproductbyscancodeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcGetproductbyscancodeAPIRequest) Reset() {
	r._code = ""
	r._codeType = ""
	r._shopCode = ""
	r._storeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.getproductbyscancode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 码, 对应的信息可能是款号，也有可能是具体的某一个商品
func (r *AlibabaMjOcGetproductbyscancodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetCode() string {
	return r._code
}

// SetCodeType is CodeType Setter
// 条码/二维码/rfid(电子标签),货号、条码、零售+唯一码;ARTNO、BARCODE、UNIQUECODE
func (r *AlibabaMjOcGetproductbyscancodeAPIRequest) SetCodeType(_codeType string) error {
	r._codeType = _codeType
	r.Set("code_type", _codeType)
	return nil
}

// GetCodeType CodeType Getter
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetCodeType() string {
	return r._codeType
}

// SetShopCode is ShopCode Setter
// 专柜编码
func (r *AlibabaMjOcGetproductbyscancodeAPIRequest) SetShopCode(_shopCode string) error {
	r._shopCode = _shopCode
	r.Set("shop_code", _shopCode)
	return nil
}

// GetShopCode ShopCode Getter
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetShopCode() string {
	return r._shopCode
}

// SetStoreCode is StoreCode Setter
// 门店编码
func (r *AlibabaMjOcGetproductbyscancodeAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaMjOcGetproductbyscancodeAPIRequest) GetStoreCode() string {
	return r._storeCode
}

var poolAlibabaMjOcGetproductbyscancodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcGetproductbyscancodeRequest()
	},
}

// GetAlibabaMjOcGetproductbyscancodeRequest 从 sync.Pool 获取 AlibabaMjOcGetproductbyscancodeAPIRequest
func GetAlibabaMjOcGetproductbyscancodeAPIRequest() *AlibabaMjOcGetproductbyscancodeAPIRequest {
	return poolAlibabaMjOcGetproductbyscancodeAPIRequest.Get().(*AlibabaMjOcGetproductbyscancodeAPIRequest)
}

// ReleaseAlibabaMjOcGetproductbyscancodeAPIRequest 将 AlibabaMjOcGetproductbyscancodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcGetproductbyscancodeAPIRequest(v *AlibabaMjOcGetproductbyscancodeAPIRequest) {
	v.Reset()
	poolAlibabaMjOcGetproductbyscancodeAPIRequest.Put(v)
}
