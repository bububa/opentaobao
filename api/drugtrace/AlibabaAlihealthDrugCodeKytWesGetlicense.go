package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytWesGetlicense 获取licenseToken
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
func AlibabaAlihealthDrugCodeKytWesGetlicense(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
