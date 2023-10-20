package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQuerydruginfo 码查询药品
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
func AlibabaAlihealthDrugKytQuerydruginfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQuerydruginfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
