package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundDetailGet 退款详情页渲染
// taobao.refund.detail.get
//
// 退款详情页渲染
func TaobaoRefundDetailGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundDetailGetAPIRequest, session string) (*tbrefund.TaobaoRefundDetailGetAPIResponse, error) {
	var resp tbrefund.TaobaoRefundDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
