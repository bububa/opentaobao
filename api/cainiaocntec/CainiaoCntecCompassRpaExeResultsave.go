package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecCompassRpaExeResultsave rpa执行结果回传
// cainiao.cntec.compass.rpa.exe.resultsave
//
// rpa执行结果回传
func CainiaoCntecCompassRpaExeResultsave(clt *core.SDKClient, req *cainiaocntec.CainiaoCntecCompassRpaExeResultsaveAPIRequest, session string) (*cainiaocntec.CainiaoCntecCompassRpaExeResultsaveAPIResponse, error) {
	var resp cainiaocntec.CainiaoCntecCompassRpaExeResultsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
