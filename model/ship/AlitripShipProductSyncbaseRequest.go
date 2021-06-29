package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
基础信息修改回调 API请求
alitrip.ship.product.syncbase

基础信息修改回调
*/
type AlitripShipProductSyncbaseRequest struct {
    model.Params
}

// 初始化AlitripShipProductSyncbaseRequest对象
func NewAlitripShipProductSyncbaseRequest() *AlitripShipProductSyncbaseRequest{
    return &AlitripShipProductSyncbaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripShipProductSyncbaseRequest) GetApiMethodName() string {
    return "alitrip.ship.product.syncbase"
}

// IRequest interface 方法, 获取API参数
func (r AlitripShipProductSyncbaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
