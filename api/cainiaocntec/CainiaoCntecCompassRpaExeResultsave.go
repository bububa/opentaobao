package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// Cainiaocnteccompassrpaexeresultsave rpa执行结果回传
// cainiao.cntec.compass.rpa.exe.resultsave
//
// rpa执行结果回传
func Cainiaocnteccompassrpaexeresultsave(clt *core.SDKClient, req *cainiaocntec.CainiaocnteccompassrpaexeresultsaveAPIRequest, session string) (*cainiaocntec.CainiaocnteccompassrpaexeresultsaveAPIResponse, error) {
	var resp cainiaocntec.CainiaocnteccompassrpaexeresultsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
