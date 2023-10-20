package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerycodeactive 查询码是否激活
// alibaba.alihealth.drug.kyt.querycodeactive
//
// 查询码是否激活
func AlibabaAlihealthDrugKytQuerycodeactive(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerycodeactiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
