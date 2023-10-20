package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductquantityupdate 产品库存修改
// taobao.fenxiao.product.quantity.update
//
// 修改产品库存信息，支持全量修改以及增量修改两种方式
func Taobaofenxiaoproductquantityupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductquantityupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductquantityupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductquantityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
