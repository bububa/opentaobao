package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeListentparbyrefentid 根据企业id获取往来单位
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
func AlibabaAlihealthDrugcodeListentparbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
