package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflow 码流向查询
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
func AlibabaAlihealthDrugCodeKytQuerycodeflow(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
