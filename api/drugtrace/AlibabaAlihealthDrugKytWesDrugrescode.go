package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesDrugrescode 查询药品码段信息
// alibaba.alihealth.drug.kyt.wes.drugrescode
//
// 查询药品码段信息
func AlibabaAlihealthDrugKytWesDrugrescode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesDrugrescodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
