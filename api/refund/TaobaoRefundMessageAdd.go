package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

/* TaobaoRefundMessageAdd
创建退款留言/凭证
taobao.refund.message.add

创建退款留言/凭证 */
func TaobaoRefundMessageAdd(clt *core.SDKClient, req *refund.TaobaoRefundMessageAddAPIRequest, session string) (*refund.TaobaoRefundMessageAddAPIResponse, error) {
	var resp refund.TaobaoRefundMessageAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
