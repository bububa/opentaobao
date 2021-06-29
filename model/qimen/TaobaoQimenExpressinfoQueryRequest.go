package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
配送公司信息查询接口 API请求
taobao.qimen.expressinfo.query

配送公司信息查询
*/
type TaobaoQimenExpressinfoQueryRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenExpressinfoQueryRequest对象
func NewTaobaoQimenExpressinfoQueryRequest() *TaobaoQimenExpressinfoQueryRequest{
    return &TaobaoQimenExpressinfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenExpressinfoQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.expressinfo.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenExpressinfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenExpressinfoQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenExpressinfoQueryRequest) GetRequest() *Request {
    return r.request
}
