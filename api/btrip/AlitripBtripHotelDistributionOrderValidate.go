package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
商旅酒店API分销下单前校验 
alitrip.btrip.hotel.distribution.order.validate

商旅酒店API分销下单前校验
*/
func AlitripBtripHotelDistributionOrderValidate(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderValidateAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionOrderValidateAPIResponse, error) {
    var resp btrip.AlitripBtripHotelDistributionOrderValidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
