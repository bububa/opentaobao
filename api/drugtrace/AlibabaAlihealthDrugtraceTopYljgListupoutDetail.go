package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljglistupoutdetail 上游出库单单据明细查询
// alibaba.alihealth.drugtrace.top.yljg.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
func Alibabaalihealthdrugtracetopyljglistupoutdetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljglistupoutdetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljglistupoutdetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljglistupoutdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
