package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkusGet SKU查询接口
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
func TaobaoFenxiaoProductSkusGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkusGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductSkusGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductSkusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
