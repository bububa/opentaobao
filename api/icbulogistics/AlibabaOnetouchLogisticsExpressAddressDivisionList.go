package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
四级地址库-区域 
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区
*/
func AlibabaOnetouchLogisticsExpressAddressDivisionList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
