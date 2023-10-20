package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodereplace 单码替换
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
func AlibabaAlihealthDrugKytCodereplace(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
