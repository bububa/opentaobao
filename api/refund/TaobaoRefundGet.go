package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

/* TaobaoRefundGet
获取单笔退款详情
taobao.refund.get

获取单笔退款详情 */
func TaobaoRefundGet(clt *core.SDKClient, req *refund.TaobaoRefundGetAPIRequest, session string) (*refund.TaobaoRefundGetAPIResponse, error) {
	var resp refund.TaobaoRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
