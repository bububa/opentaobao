package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyDelbillinfo 根据单据号删除单据
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
func AlibabaAlihealthDrugKytScqyDelbillinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyDelbillinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
