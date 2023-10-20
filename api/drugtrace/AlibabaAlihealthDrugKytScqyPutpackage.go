package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyPutpackage 码拼箱
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
func AlibabaAlihealthDrugKytScqyPutpackage(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
