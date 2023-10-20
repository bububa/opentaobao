package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugLsydSaveent 零售药店往来单位新增
// alibaba.alihealth.drug.lsyd.saveent
//
// 新增往来单位企业记录
func AlibabaAlihealthDrugLsydSaveent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugLsydSaveentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugLsydSaveentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
