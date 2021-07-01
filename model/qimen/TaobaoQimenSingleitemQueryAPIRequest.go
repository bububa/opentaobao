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
type TaobaoQimenSingleitemQueryAPIRequest struct {
    model.Params
    // 
    _request   *RequestDo
}

// 初始化TaobaoQimenSingleitemQueryAPIRequest对象
func NewTaobaoQimenSingleitemQueryRequest() *TaobaoQimenSingleitemQueryAPIRequest{
    return &TaobaoQimenSingleitemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemQueryAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.singleitem.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenSingleitemQueryAPIRequest) SetRequest(_request *RequestDo) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenSingleitemQueryAPIRequest) GetRequest() *RequestDo {
    return r._request
}
