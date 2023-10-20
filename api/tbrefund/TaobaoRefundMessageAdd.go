package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundmessageadd 创建退款留言/凭证
// taobao.refund.message.add
//
// 创建退款留言/凭证
func Taobaorefundmessageadd(clt *core.SDKClient, req *tbrefund.TaobaorefundmessageaddAPIRequest, session string) (*tbrefund.TaobaorefundmessageaddAPIResponse, error) {
	var resp tbrefund.TaobaorefundmessageaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
