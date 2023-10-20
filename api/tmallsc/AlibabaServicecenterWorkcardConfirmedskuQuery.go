package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardConfirmedskuQuery 查询确认履行的服务项
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
func AlibabaServicecenterWorkcardConfirmedskuQuery(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest, session string) (*tmallsc.AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse, error) {
	var resp tmallsc.AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
