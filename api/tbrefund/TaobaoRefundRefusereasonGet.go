package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundRefusereasonGet 获取拒绝原因列表
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
func TaobaoRefundRefusereasonGet(clt *core.SDKClient, req *tbrefund.TaobaoRefundRefusereasonGetAPIRequest, session string) (*tbrefund.TaobaoRefundRefusereasonGetAPIResponse, error) {
	var resp tbrefund.TaobaoRefundRefusereasonGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
