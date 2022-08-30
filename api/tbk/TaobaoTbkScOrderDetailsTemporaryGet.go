package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScOrderDetailsTemporaryGet 淘宝客-服务商-所有订单查询（临时接口）
// taobao.tbk.sc.order.details.temporary.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
func TaobaoTbkScOrderDetailsTemporaryGet(clt *core.SDKClient, req *tbk.TaobaoTbkScOrderDetailsTemporaryGetAPIRequest, session string) (*tbk.TaobaoTbkScOrderDetailsTemporaryGetAPIResponse, error) {
	var resp tbk.TaobaoTbkScOrderDetailsTemporaryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
