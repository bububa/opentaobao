package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytlistupout 查询货主/本企业上游企业出库单据信息
// alibaba.alihealth.drug.kyt.listupout
//
// 查询货主/本企业上游企业出库单据信息
func Alibabaalihealthdrugkytlistupout(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytlistupoutAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytlistupoutAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytlistupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
