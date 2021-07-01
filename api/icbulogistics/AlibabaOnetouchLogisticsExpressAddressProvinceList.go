package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
四级地址库-省 
alibaba.onetouch.logistics.express.address.province.list

四级地址库-省
*/
func AlibabaOnetouchLogisticsExpressAddressProvinceList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
