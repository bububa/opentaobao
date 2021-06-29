package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
四级地址库-市 
alibaba.onetouch.logistics.express.address.city.list

四级地址库-市
*/
func AlibabaOnetouchLogisticsExpressAddressCityList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
