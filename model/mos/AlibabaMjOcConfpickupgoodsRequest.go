package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提货核销 API请求
alibaba.mj.oc.confpickupgoods

此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
*/
type AlibabaMjOcConfpickupgoodsRequest struct {
    model.Params
    // 提货核销请求参数
    confPickupGoodsRequest   *ConfPickupGoodsReqDto
}

// 初始化AlibabaMjOcConfpickupgoodsRequest对象
func NewAlibabaMjOcConfpickupgoodsRequest() *AlibabaMjOcConfpickupgoodsRequest{
    return &AlibabaMjOcConfpickupgoodsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcConfpickupgoodsRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.confpickupgoods"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcConfpickupgoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConfPickupGoodsRequest Setter
// 提货核销请求参数
func (r *AlibabaMjOcConfpickupgoodsRequest) SetConfPickupGoodsRequest(confPickupGoodsRequest *ConfPickupGoodsReqDto) error {
    r.confPickupGoodsRequest = confPickupGoodsRequest
    r.Set("conf_pickup_goods_request", confPickupGoodsRequest)
    return nil
}

// ConfPickupGoodsRequest Getter
func (r AlibabaMjOcConfpickupgoodsRequest) GetConfPickupGoodsRequest() *ConfPickupGoodsReqDto {
    return r.confPickupGoodsRequest
}
