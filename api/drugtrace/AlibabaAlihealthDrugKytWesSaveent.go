package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSaveent 新增往来单位企业记录
// alibaba.alihealth.drug.kyt.wes.saveent
//
// 新增往来单位企业记录
func AlibabaAlihealthDrugKytWesSaveent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSaveentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSaveentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
