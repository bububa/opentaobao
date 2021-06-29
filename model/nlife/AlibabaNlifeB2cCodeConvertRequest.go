package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c转码 APIRequest
alibaba.nlife.b2c.code.convert

将商品的URL转码，ISV将该码写入RFID
*/
type AlibabaNlifeB2cCodeConvertRequest struct {
    model.Params

    // 零售商在零售+平台ID，非唯一码模式必填，建议传递该值
    storeId   string 

    // 商品URL
    url   string 

}

func NewAlibabaNlifeB2cCodeConvertRequest() *AlibabaNlifeB2cCodeConvertRequest{
    return &AlibabaNlifeB2cCodeConvertRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeB2cCodeConvertRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.code.convert"
}

func (r AlibabaNlifeB2cCodeConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeB2cCodeConvertRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeB2cCodeConvertRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaNlifeB2cCodeConvertRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaNlifeB2cCodeConvertRequest) GetUrl() string {
    return r.url
}

