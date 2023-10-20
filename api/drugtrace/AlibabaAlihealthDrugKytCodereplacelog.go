package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodereplacelog 码替换记录查询
// alibaba.alihealth.drug.kyt.codereplacelog
//
// 码替换记录查询
func AlibabaAlihealthDrugKytCodereplacelog(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodereplacelogAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytCodereplacelogAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
