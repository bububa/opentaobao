package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoSpecialRefundGet 特殊部分退纠纷单查询
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
func TaobaoSpecialRefundGet(clt *core.SDKClient, req *tbrefund.TaobaoSpecialRefundGetAPIRequest, session string) (*tbrefund.TaobaoSpecialRefundGetAPIResponse, error) {
	var resp tbrefund.TaobaoSpecialRefundGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
