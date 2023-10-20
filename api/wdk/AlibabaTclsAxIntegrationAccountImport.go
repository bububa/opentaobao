package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAxIntegrationAccountImport ISV用户录入
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
func AlibabaTclsAxIntegrationAccountImport(clt *core.SDKClient, req *wdk.AlibabaTclsAxIntegrationAccountImportAPIRequest, resp *wdk.AlibabaTclsAxIntegrationAccountImportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
