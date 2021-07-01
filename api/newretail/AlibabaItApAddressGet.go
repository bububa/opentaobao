package newretail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/newretail"
)

/* 
getApAddressByMacNew 
alibaba.it.ap.address.get

根据ap 的mac地址查询ap的结构化位置信息
*/
func AlibabaItApAddressGet(clt *core.SDKClient, req *newretail.AlibabaItApAddressGetAPIRequest, session string) (*newretail.AlibabaItApAddressGetAPIResponse, error) {
    var resp newretail.AlibabaItApAddressGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
