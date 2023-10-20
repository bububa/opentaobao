package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundGet 获取单笔退款详情
// taobao.refund.get
//
// 获取单笔退款详情
func TaobaoRefundGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundGetAPIRequest, session string) (*tbrefund.TaobaoRefundGetAPIResponse, error) {
	var resp tbrefund.TaobaoRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
