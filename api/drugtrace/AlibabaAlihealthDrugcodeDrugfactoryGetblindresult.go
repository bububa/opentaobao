package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresult 获取盲底文件处理结果
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
func AlibabaAlihealthDrugcodeDrugfactoryGetblindresult(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
