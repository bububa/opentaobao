package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaxintegrationaccountimport ISV用户录入
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
func Alibabatclsaxintegrationaccountimport(clt *core.SDKClient, req *wdk.AlibabatclsaxintegrationaccountimportAPIRequest, session string) (*wdk.AlibabatclsaxintegrationaccountimportAPIResponse, error) {
	var resp wdk.AlibabatclsaxintegrationaccountimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
