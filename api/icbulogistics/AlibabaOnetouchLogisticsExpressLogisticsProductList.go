package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// Alibabaonetouchlogisticsexpresslogisticsproductlist 查询物流运力列表
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
func Alibabaonetouchlogisticsexpresslogisticsproductlist(clt *core.SDKClient, req *icbulogistics.AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest, session string) (*icbulogistics.AlibabaonetouchlogisticsexpresslogisticsproductlistAPIResponse, error) {
	var resp icbulogistics.AlibabaonetouchlogisticsexpresslogisticsproductlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
