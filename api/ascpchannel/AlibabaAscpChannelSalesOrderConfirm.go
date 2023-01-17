package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSalesOrderConfirm 渠道销售单确认收货
// alibaba.ascp.channel.sales.order.confirm
//
// 渠道销售单确认收货
func AlibabaAscpChannelSalesOrderConfirm(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSalesOrderConfirmAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSalesOrderConfirmAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSalesOrderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
