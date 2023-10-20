package icbulogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressLogisticsProductList 查询物流运力列表
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
func AlibabaOnetouchLogisticsExpressLogisticsProductList(clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressLogisticsProductListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
