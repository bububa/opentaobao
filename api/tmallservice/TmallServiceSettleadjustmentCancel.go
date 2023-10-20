package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicesettleadjustmentcancel 取消结算调整单
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
func Tmallservicesettleadjustmentcancel(clt *core.SDKClient, req *tmallservice.TmallservicesettleadjustmentcancelAPIRequest, session string) (*tmallservice.TmallservicesettleadjustmentcancelAPIResponse, error) {
	var resp tmallservice.TmallservicesettleadjustmentcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
