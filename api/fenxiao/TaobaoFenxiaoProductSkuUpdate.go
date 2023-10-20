package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductskuupdate 产品sku编辑接口
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
func Taobaofenxiaoproductskuupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductskuupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductskuupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
