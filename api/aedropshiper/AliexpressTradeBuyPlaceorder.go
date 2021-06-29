package aedropshiper

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aedropshiper"
)

/* 
AE下单API 
aliexpress.trade.buy.placeorder

A006_INVALID_ACCOUNT_INFO
*/
func AliexpressTradeBuyPlaceorder(clt *core.SDKClient, req *aedropshiper.AliexpressTradeBuyPlaceorderRequest, session string) (*aedropshiper.AliexpressTradeBuyPlaceorderAPIResponse, error) {
    var resp aedropshiper.AliexpressTradeBuyPlaceorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
