package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
淘外分销逆向创单（未发货整单退） 
alibaba.ascp.channel.main.refund.create

淘外分销解决方案--订单--逆向创单（未发货整单退）
*/
func AlibabaAscpChannelMainRefundCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelMainRefundCreateAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelMainRefundCreateAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpChannelMainRefundCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
