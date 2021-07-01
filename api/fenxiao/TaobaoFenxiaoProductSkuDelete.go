package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoProductSkuDelete
产品SKU删除接口
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据 */
func TaobaoFenxiaoProductSkuDelete(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuDeleteAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductSkuDeleteAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductSkuDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
