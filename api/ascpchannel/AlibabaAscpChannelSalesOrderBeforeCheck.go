package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelSalesOrderBeforeCheck 供应链外部订单创建前校验接口
// alibaba.ascp.channel.sales.order.before.check
//
// 供应链外部订单创建前校验接口
func AlibabaAscpChannelSalesOrderBeforeCheck(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelSalesOrderBeforeCheckAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelSalesOrderBeforeCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
