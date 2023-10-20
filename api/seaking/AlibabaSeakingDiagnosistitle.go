package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingdiagnosistitle 标题诊断
// alibaba.seaking.diagnosistitle
//
// 标题诊断
func Alibabaseakingdiagnosistitle(clt *core.SDKClient, req *seaking.AlibabaseakingdiagnosistitleAPIRequest, session string) (*seaking.AlibabaseakingdiagnosistitleAPIResponse, error) {
	var resp seaking.AlibabaseakingdiagnosistitleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
