package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductskudelete 产品SKU删除接口
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
func Taobaofenxiaoproductskudelete(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductskudeleteAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductskudeleteAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductskudeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
