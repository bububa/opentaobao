package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfo 获取盲底文件中的批次信息
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
func AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
