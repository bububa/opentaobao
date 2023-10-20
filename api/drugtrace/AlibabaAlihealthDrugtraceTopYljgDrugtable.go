package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgdrugtable 查询药品目录信息
// alibaba.alihealth.drugtrace.top.yljg.drugtable
//
// 查询药品目录信息
func Alibabaalihealthdrugtracetopyljgdrugtable(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgdrugtableAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgdrugtableAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgdrugtableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
