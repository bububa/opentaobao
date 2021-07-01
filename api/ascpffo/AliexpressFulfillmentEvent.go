package ascpffo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpffo"
)

/* 
AE履约事件处理 
aliexpress.fulfillment.event

AE用 履约底层声明发货能力
*/
func AliexpressFulfillmentEvent(clt *core.SDKClient, req *ascpffo.AliexpressFulfillmentEventAPIRequest, session string) (*ascpffo.AliexpressFulfillmentEventAPIResponse, error) {
    var resp ascpffo.AliexpressFulfillmentEventAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
