package icbulogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbulogistics"
)

/* 
计算快递运费&下单参数校验 
alibaba.onetouch.logistics.express.charge.calculate

计算快递运费、下单参数校验
*/
func AlibabaOnetouchLogisticsExpressChargeCalculate(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest, session string) (*icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse, error) {
    var resp icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
