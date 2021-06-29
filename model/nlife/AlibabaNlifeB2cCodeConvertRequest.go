package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c转码 API请求
alibaba.nlife.b2c.code.convert

将商品的URL转码，ISV将该码写入RFID
*/
type AlibabaNlifeB2cCodeConvertRequest struct {
    model.Params
    // 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
    _storeId   string
    // 商品URL
    _url   string
}

// 初始化AlibabaNlifeB2cCodeConvertRequest对象
func NewAlibabaNlifeB2cCodeConvertRequest() *AlibabaNlifeB2cCodeConvertRequest{
    return &AlibabaNlifeB2cCodeConvertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cCodeConvertRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.code.convert"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cCodeConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
func (r *AlibabaNlifeB2cCodeConvertRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cCodeConvertRequest) GetStoreId() string {
    return r._storeId
}
// Url Setter
// 商品URL
func (r *AlibabaNlifeB2cCodeConvertRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaNlifeB2cCodeConvertRequest) GetUrl() string {
    return r._url
}
