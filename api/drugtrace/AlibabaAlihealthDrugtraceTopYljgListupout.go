package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljglistupout 医疗机构查询本企业上游企业出库单据信息
// alibaba.alihealth.drugtrace.top.yljg.listupout
//
// 查询货主/本企业上游企业出库单据信息
func Alibabaalihealthdrugtracetopyljglistupout(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljglistupoutAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljglistupoutAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljglistupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
