package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.drugrescode
//
// 查询药品码段信息
func AlibabaAlihealthDrugKytDrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugrescodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
