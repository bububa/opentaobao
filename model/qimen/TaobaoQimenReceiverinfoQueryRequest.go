package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID 收件人信息解密接口 API请求
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息
*/
type TaobaoQimenReceiverinfoQueryRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenReceiverinfoQueryRequest对象
func NewTaobaoQimenReceiverinfoQueryRequest() *TaobaoQimenReceiverinfoQueryRequest{
    return &TaobaoQimenReceiverinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenReceiverinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.receiverinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenReceiverinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenReceiverinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenReceiverinfoQueryRequest) GetRequest() *Request {
    return r.request
}
