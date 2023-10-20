package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetcodebillinfo 根据码获取基本和单据信息
// alibaba.alihealth.drug.kyt.getcodebillinfo
//
// 根据码信息获取基本信息和单据信息
func AlibabaAlihealthDrugKytGetcodebillinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetcodebillinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
