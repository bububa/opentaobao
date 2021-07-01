package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 API请求
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteAPIRequest struct {
    model.Params
    // 
    _request   *RequestDo
}

// 初始化TaobaoQimenCombineitemDeleteAPIRequest对象
func NewTaobaoQimenCombineitemDeleteRequest() *TaobaoQimenCombineitemDeleteAPIRequest{
    return &TaobaoQimenCombineitemDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenCombineitemDeleteAPIRequest) SetRequest(_request *RequestDo) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenCombineitemDeleteAPIRequest) GetRequest() *RequestDo {
    return r._request
}
