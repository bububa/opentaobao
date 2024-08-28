package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAxIntegrationAccountImport ISV用户录入
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
func AlibabaTclsAxIntegrationAccountImport(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAxIntegrationAccountImportAPIRequest, resp *wdk.AlibabaTclsAxIntegrationAccountImportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
