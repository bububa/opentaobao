package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytListpartsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
