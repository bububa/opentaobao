package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticssellerwritelogisticsnode 商家配送写入物流节点
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
func Alibabaascplogisticssellerwritelogisticsnode(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticssellerwritelogisticsnodeAPIRequest, session string) (*tblogistics.AlibabaascplogisticssellerwritelogisticsnodeAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticssellerwritelogisticsnodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
