package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticssellerwriteoff 商家配送核销
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
func Alibabaascplogisticssellerwriteoff(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticssellerwriteoffAPIRequest, session string) (*tblogistics.AlibabaascplogisticssellerwriteoffAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticssellerwriteoffAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
