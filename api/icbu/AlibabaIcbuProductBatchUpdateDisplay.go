package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductBatchUpdateDisplay 商品批量上下架接口
// alibaba.icbu.product.batch.update.display
//
// 给国际站的三方服务商提供批量上下架接口
func AlibabaIcbuProductBatchUpdateDisplay(clt *core.SDKClient, req *icbu.AlibabaIcbuProductBatchUpdateDisplayAPIRequest, resp *icbu.AlibabaIcbuProductBatchUpdateDisplayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
