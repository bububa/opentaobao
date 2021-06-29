package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 API请求
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenCombineitemQueryRequest对象
func NewTaobaoQimenCombineitemQueryRequest() *TaobaoQimenCombineitemQueryRequest{
    return &TaobaoQimenCombineitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenCombineitemQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenCombineitemQueryRequest) GetRequest() *Request {
    return r.request
}
