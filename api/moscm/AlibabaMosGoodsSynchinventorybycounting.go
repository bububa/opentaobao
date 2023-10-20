package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosgoodssynchinventorybycounting 以盘点方式调整库存：传入商品实际库存
// alibaba.mos.goods.synchinventorybycounting
//
// 以盘点方式调整库存：传入商品实际库存
// 盘点单自动判断数量增减
func Alibabamosgoodssynchinventorybycounting(clt *core.SDKClient, req *moscm.AlibabamosgoodssynchinventorybycountingAPIRequest, session string) (*moscm.AlibabamosgoodssynchinventorybycountingAPIResponse, error) {
	var resp moscm.AlibabamosgoodssynchinventorybycountingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
