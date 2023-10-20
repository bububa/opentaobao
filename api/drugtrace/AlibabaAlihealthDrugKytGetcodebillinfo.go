package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytgetcodebillinfo 根据码获取基本和单据信息
// alibaba.alihealth.drug.kyt.getcodebillinfo
//
// 根据码信息获取基本信息和单据信息
func Alibabaalihealthdrugkytgetcodebillinfo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytgetcodebillinfoAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytgetcodebillinfoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytgetcodebillinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
