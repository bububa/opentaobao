package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventoryadjustexternal 非交易库存调整单
// taobao.inventory.adjust.external
//
// 商家非交易调整库存，调拨出库、盘点等时调用
func Taobaoinventoryadjustexternal(clt *core.SDKClient, req *fenxiao.TaobaoinventoryadjustexternalAPIRequest, session string) (*fenxiao.TaobaoinventoryadjustexternalAPIResponse, error) {
	var resp fenxiao.TaobaoinventoryadjustexternalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
