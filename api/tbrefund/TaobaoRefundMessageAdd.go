package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundMessageAdd 创建退款留言/凭证
// taobao.refund.message.add
//
// 创建退款留言/凭证
func TaobaoRefundMessageAdd(clt *core.SDKClient, req *tbrefund.TaobaoRefundMessageAddAPIRequest, session string) (*tbrefund.TaobaoRefundMessageAddAPIResponse, error) {
	var resp tbrefund.TaobaoRefundMessageAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
