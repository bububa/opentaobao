package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkuAdd 产品sku添加接口
// taobao.fenxiao.product.sku.add
//
// 添加产品SKU信息
func TaobaoFenxiaoProductSkuAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuAddAPIRequest, resp *fenxiao.TaobaoFenxiaoProductSkuAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
