package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoProductSkuAdd
产品sku添加接口
taobao.fenxiao.product.sku.add

添加产品SKU信息 */
func TaobaoFenxiaoProductSkuAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuAddAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductSkuAddAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductSkuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
