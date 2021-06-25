package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID 收件人信息解密接口 APIRequest
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息
*/
type TaobaoQimenReceiverinfoQueryRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenReceiverinfoQueryRequest() *TaobaoQimenReceiverinfoQueryRequest{
    return &TaobaoQimenReceiverinfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenReceiverinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.receiverinfo.query"
}

func (r TaobaoQimenReceiverinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenReceiverinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenReceiverinfoQueryRequest) GetRequest() *Request {
    return r.request
}

