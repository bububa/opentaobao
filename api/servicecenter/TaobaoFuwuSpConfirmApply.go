package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuSpConfirmApply 内购服务确认单申请接口
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
func TaobaoFuwuSpConfirmApply(clt *core.SDKClient, req *servicecenter.TaobaoFuwuSpConfirmApplyAPIRequest, session string) (*servicecenter.TaobaoFuwuSpConfirmApplyAPIResponse, error) {
	var resp servicecenter.TaobaoFuwuSpConfirmApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
