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
type TaobaoQimenCombineitemSynchronizeAPIRequest struct {
    model.Params
    // 
    _request   *CombineItemSyncRequest
}

// 初始化TaobaoQimenCombineitemSynchronizeAPIRequest对象
func NewTaobaoQimenCombineitemSynchronizeRequest() *TaobaoQimenCombineitemSynchronizeAPIRequest{
    return &TaobaoQimenCombineitemSynchronizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.combineitem.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenCombineitemSynchronizeAPIRequest) SetRequest(_request *CombineItemSyncRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetRequest() *CombineItemSyncRequest {
    return r._request
}
