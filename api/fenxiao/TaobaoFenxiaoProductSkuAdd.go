package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductskuadd 产品sku添加接口
// taobao.fenxiao.product.sku.add
//
// 添加产品SKU信息
func Taobaofenxiaoproductskuadd(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductskuaddAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductskuaddAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
