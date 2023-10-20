package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytweslistupout 查询货主/本企业上游企业出库单据信息
// alibaba.alihealth.drug.kyt.wes.listupout
//
// 查询货主/本企业上游企业出库单据信息
func Alibabaalihealthdrugkytweslistupout(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytweslistupoutAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytweslistupoutAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytweslistupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
