package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsSellerWriteoff 商家配送核销
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
func AlibabaAscpLogisticsSellerWriteoff(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsSellerWriteoffAPIRequest, session string) (*tblogistics.AlibabaAscpLogisticsSellerWriteoffAPIResponse, error) {
	var resp tblogistics.AlibabaAscpLogisticsSellerWriteoffAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
