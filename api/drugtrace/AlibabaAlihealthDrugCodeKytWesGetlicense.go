package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytWesGetlicense 获取licenseToken
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
func AlibabaAlihealthDrugCodeKytWesGetlicense(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytWesGetlicenseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
