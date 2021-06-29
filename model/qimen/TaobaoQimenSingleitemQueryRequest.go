package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 API请求
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryRequest struct {
    model.Params
    // 
    request   *RequestDO
}

// 初始化TaobaoQimenSingleitemQueryRequest对象
func NewTaobaoQimenSingleitemQueryRequest() *TaobaoQimenSingleitemQueryRequest{
    return &TaobaoQimenSingleitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.singleitem.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenSingleitemQueryRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenSingleitemQueryRequest) GetRequest() *RequestDO {
    return r.request
}
