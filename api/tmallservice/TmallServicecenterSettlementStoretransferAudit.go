package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentersettlementstoretransferaudit 新康众审批门店分账
// tmall.servicecenter.settlement.storetransfer.audit
//
// 新康众审批门店分账
func Tmallservicecentersettlementstoretransferaudit(clt *core.SDKClient, req *tmallservice.TmallservicecentersettlementstoretransferauditAPIRequest, session string) (*tmallservice.TmallservicecentersettlementstoretransferauditAPIResponse, error) {
	var resp tmallservice.TmallservicecentersettlementstoretransferauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
