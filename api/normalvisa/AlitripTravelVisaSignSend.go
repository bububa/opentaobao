package normalvisa

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/normalvisa"
)

/* 
签证批量申请人送签接口 
alitrip.travel.visa.sign.send

签证批量申请人送签接口，用于商家批量送签。
*/
func AlitripTravelVisaSignSend(clt *core.SDKClient, req *normalvisa.AlitripTravelVisaSignSendRequest, session string) (*normalvisa.AlitripTravelVisaSignSendAPIResponse, error) {
    var resp normalvisa.AlitripTravelVisaSignSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
