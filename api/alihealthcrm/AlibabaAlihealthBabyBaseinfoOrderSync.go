package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthBabyBaseinfoOrderSync alibaba.alihealth.baby.baseinfo.order.sync
// alibaba.alihealth.baby.baseinfo.order.sync
//
// 育学园将订单信息回传给我们
func AlibabaAlihealthBabyBaseinfoOrderSync(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
