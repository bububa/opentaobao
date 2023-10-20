package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingDiagnosistitle 标题诊断
// alibaba.seaking.diagnosistitle
//
// 标题诊断
func AlibabaSeakingDiagnosistitle(clt *core.SDKClient, req *seaking.AlibabaSeakingDiagnosistitleAPIRequest, resp *seaking.AlibabaSeakingDiagnosistitleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
