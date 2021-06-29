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
type TaobaoQimenCombineitemDeleteRequest struct {
    model.Params
    // 
    request   *RequestDO
}

// 初始化TaobaoQimenCombineitemDeleteRequest对象
func NewTaobaoQimenCombineitemDeleteRequest() *TaobaoQimenCombineitemDeleteRequest{
    return &TaobaoQimenCombineitemDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenCombineitemDeleteRequest) SetRequest(request *RequestDO) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenCombineitemDeleteRequest) GetRequest() *RequestDO {
    return r.request
}
