package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaospecialrefundsreceiveget 特殊退款类型的纠纷单列表查询
// taobao.special.refunds.receive.get
//
// 特殊退款类型的纠纷单列表查询
func Taobaospecialrefundsreceiveget(clt *core.SDKClient, req *tbrefund.TaobaospecialrefundsreceivegetAPIRequest, session string) (*tbrefund.TaobaospecialrefundsreceivegetAPIResponse, error) {
	var resp tbrefund.TaobaospecialrefundsreceivegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
