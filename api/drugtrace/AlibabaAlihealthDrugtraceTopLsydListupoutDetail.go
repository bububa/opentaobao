package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetoplsydlistupoutdetail 上游出库单单据明细查询
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
func Alibabaalihealthdrugtracetoplsydlistupoutdetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
