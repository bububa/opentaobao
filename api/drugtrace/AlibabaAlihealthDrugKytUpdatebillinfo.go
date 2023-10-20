package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpdatebillinfo 零售端平台单据更新
// alibaba.alihealth.drug.kyt.updatebillinfo
//
// 零售端平台单据更新
func AlibabaAlihealthDrugKytUpdatebillinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
