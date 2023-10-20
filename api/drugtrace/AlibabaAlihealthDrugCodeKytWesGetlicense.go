package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekytwesgetlicense 获取licenseToken
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
func Alibabaalihealthdrugcodekytwesgetlicense(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekytwesgetlicenseAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekytwesgetlicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
