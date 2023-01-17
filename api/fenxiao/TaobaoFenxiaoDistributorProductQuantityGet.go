package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorProductQuantityGet 分销商查询产品库存
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
func TaobaoFenxiaoDistributorProductQuantityGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoDistributorProductQuantityGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
