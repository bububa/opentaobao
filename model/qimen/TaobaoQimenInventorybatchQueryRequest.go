package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品单仓批次库存查询接口 API请求
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存
*/
type TaobaoQimenInventorybatchQueryRequest struct {
    model.Params
    // request
    request   *Request
}

// 初始化TaobaoQimenInventorybatchQueryRequest对象
func NewTaobaoQimenInventorybatchQueryRequest() *TaobaoQimenInventorybatchQueryRequest{
    return &TaobaoQimenInventorybatchQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorybatchQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.inventorybatch.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorybatchQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// request
func (r *TaobaoQimenInventorybatchQueryRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenInventorybatchQueryRequest) GetRequest() *Request {
    return r.request
}
