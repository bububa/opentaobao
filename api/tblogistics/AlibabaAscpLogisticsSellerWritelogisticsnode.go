package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerWritelogisticsnode 商家配送写入物流节点
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
func AlibabaAscpLogisticsSellerWritelogisticsnode(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
