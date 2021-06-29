package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全量同步回调 API请求
alitrip.ship.product.synccall

全量同步接口
*/
type AlitripShipProductSynccallRequest struct {
    model.Params
}

// 初始化AlitripShipProductSynccallRequest对象
func NewAlitripShipProductSynccallRequest() *AlitripShipProductSynccallRequest{
    return &AlitripShipProductSynccallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripShipProductSynccallRequest) GetApiMethodName() string {
    return "alitrip.ship.product.synccall"
}

// IRequest interface 方法, 获取API参数
func (r AlitripShipProductSynccallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
