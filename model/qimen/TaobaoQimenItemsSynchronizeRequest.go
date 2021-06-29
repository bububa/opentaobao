package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 (批量) API请求
taobao.qimen.items.synchronize

ERP调用奇门的接口,批量同步商品信息给WMS
*/
type TaobaoQimenItemsSynchronizeRequest struct {
    model.Params
    // 
    request   *ItemsSynchronizeRequest
}

// 初始化TaobaoQimenItemsSynchronizeRequest对象
func NewTaobaoQimenItemsSynchronizeRequest() *TaobaoQimenItemsSynchronizeRequest{
    return &TaobaoQimenItemsSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.items.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenItemsSynchronizeRequest) SetRequest(request *ItemsSynchronizeRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenItemsSynchronizeRequest) GetRequest() *ItemsSynchronizeRequest {
    return r.request
}
