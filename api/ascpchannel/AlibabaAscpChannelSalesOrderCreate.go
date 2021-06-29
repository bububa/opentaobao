package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
供应链渠道销售单创建接口 
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口
*/
func AlibabaAscpChannelSalesOrderCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSalesOrderCreateRequest, session string) (*ascpchannel.AlibabaAscpChannelSalesOrderCreateAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpChannelSalesOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
