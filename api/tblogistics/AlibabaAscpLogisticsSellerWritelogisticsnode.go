package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerWritelogisticsnode 商家配送写入物流节点
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
func AlibabaAscpLogisticsSellerWritelogisticsnode(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest, resp *tblogistics.AlibabaAscpLogisticsSellerWritelogisticsnodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
