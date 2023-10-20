package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerWriteoff 商家配送核销
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
func AlibabaAscpLogisticsSellerWriteoff(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerWriteoffAPIRequest, resp *tblogistics.AlibabaAscpLogisticsSellerWriteoffAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
