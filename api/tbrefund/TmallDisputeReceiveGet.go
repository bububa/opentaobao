package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TmallDisputeReceiveGet 天猫逆向纠纷查询
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
func TmallDisputeReceiveGet(clt *core.SDKClient, req *tbrefund.TmallDisputeReceiveGetAPIRequest, session string) (*tbrefund.TmallDisputeReceiveGetAPIResponse, error) {
	var resp tbrefund.TmallDisputeReceiveGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
