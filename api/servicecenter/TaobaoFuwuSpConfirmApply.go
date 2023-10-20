package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaofuwuspconfirmapply 内购服务确认单申请接口
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
func Taobaofuwuspconfirmapply(clt *core.SDKClient, req *servicecenter.TaobaofuwuspconfirmapplyAPIRequest, session string) (*servicecenter.TaobaofuwuspconfirmapplyAPIResponse, error) {
	var resp servicecenter.TaobaofuwuspconfirmapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
