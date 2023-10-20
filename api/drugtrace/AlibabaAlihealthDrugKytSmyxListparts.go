package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSmyxListparts 药店查询往来单位
// alibaba.alihealth.drug.kyt.smyx.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytSmyxListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSmyxListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytSmyxListpartsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
