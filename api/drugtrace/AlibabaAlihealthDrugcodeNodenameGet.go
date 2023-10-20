package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeNodenameGet 根据码获取机构名称
// alibaba.alihealth.drugcode.nodename.get
//
// 根据码获取机构名称
func AlibabaAlihealthDrugcodeNodenameGet(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeNodenameGetAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeNodenameGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
