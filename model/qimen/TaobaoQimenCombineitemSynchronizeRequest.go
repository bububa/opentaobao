package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品接口 API请求
taobao.qimen.combineitem.synchronize

ERP调用奇门的接口,将商品信息同步给WMS
*/
type TaobaoQimenCombineitemSynchronizeRequest struct {
    model.Params
    // 
    _request   *CombineItemSyncRequest
}

// 初始化TaobaoQimenCombineitemSynchronizeRequest对象
func NewTaobaoQimenCombineitemSynchronizeRequest() *TaobaoQimenCombineitemSynchronizeRequest{
    return &TaobaoQimenCombineitemSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenCombineitemSynchronizeRequest) SetRequest(_request *CombineItemSyncRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenCombineitemSynchronizeRequest) GetRequest() *CombineItemSyncRequest {
    return r._request
}
