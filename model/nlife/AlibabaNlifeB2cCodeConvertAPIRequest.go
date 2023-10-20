package nlife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cCodeConvertAPIRequest b2c转码 API请求
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
type AlibabaNlifeB2cCodeConvertAPIRequest struct {
	model.Params
	// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
	_storeId string
	// 商品URL
	_url string
}

// NewAlibabaNlifeB2cCodeConvertRequest 初始化AlibabaNlifeB2cCodeConvertAPIRequest对象
func NewAlibabaNlifeB2cCodeConvertRequest() *AlibabaNlifeB2cCodeConvertAPIRequest {
	return &AlibabaNlifeB2cCodeConvertAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cCodeConvertAPIRequest) Reset() {
	r._storeId = ""
	r._url = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cCodeConvertAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.code.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cCodeConvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cCodeConvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
func (r *AlibabaNlifeB2cCodeConvertAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cCodeConvertAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetUrl is Url Setter
// 商品URL
func (r *AlibabaNlifeB2cCodeConvertAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaNlifeB2cCodeConvertAPIRequest) GetUrl() string {
	return r._url
}

var poolAlibabaNlifeB2cCodeConvertAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cCodeConvertRequest()
	},
}

// GetAlibabaNlifeB2cCodeConvertRequest 从 sync.Pool 获取 AlibabaNlifeB2cCodeConvertAPIRequest
func GetAlibabaNlifeB2cCodeConvertAPIRequest() *AlibabaNlifeB2cCodeConvertAPIRequest {
	return poolAlibabaNlifeB2cCodeConvertAPIRequest.Get().(*AlibabaNlifeB2cCodeConvertAPIRequest)
}

// ReleaseAlibabaNlifeB2cCodeConvertAPIRequest 将 AlibabaNlifeB2cCodeConvertAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cCodeConvertAPIRequest(v *AlibabaNlifeB2cCodeConvertAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cCodeConvertAPIRequest.Put(v)
}
