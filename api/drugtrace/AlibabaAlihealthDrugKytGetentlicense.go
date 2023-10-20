package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetentlicense 获取企业资质
// alibaba.alihealth.drug.kyt.getentlicense
//
// 获取企业的资质信息。
func AlibabaAlihealthDrugKytGetentlicense(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetentlicenseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
