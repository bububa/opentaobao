package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAxIntegrationAccountImport ISV用户录入
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
func AlibabaTclsAxIntegrationAccountImport(clt *core.SDKClient, req *wdk.AlibabaTclsAxIntegrationAccountImportAPIRequest, session string) (*wdk.AlibabaTclsAxIntegrationAccountImportAPIResponse, error) {
	var resp wdk.AlibabaTclsAxIntegrationAccountImportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
