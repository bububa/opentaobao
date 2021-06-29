package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 API请求
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
type TaobaoQimenSingleitemSynchronizeRequest struct {
    model.Params
    // 
    _request   *ItemSynRequest
}

// 初始化TaobaoQimenSingleitemSynchronizeRequest对象
func NewTaobaoQimenSingleitemSynchronizeRequest() *TaobaoQimenSingleitemSynchronizeRequest{
    return &TaobaoQimenSingleitemSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.singleitem.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenSingleitemSynchronizeRequest) SetRequest(_request *ItemSynRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenSingleitemSynchronizeRequest) GetRequest() *ItemSynRequest {
    return r._request
}
