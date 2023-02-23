package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsOfflineSend 自己联系物流发货
// alibaba.ascp.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
func AlibabaAscpLogisticsOfflineSend(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsOfflineSendAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsOfflineSendAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsOfflineSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
