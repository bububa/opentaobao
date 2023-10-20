package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveState 体检机构对接_体检状态查询
// alibaba.alihealth.examination.reserve.state
//
// 体检机构对接_体检状态查询
func AlibabaAlihealthExaminationReserveState(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveStateAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveStateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
