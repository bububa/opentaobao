package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingUserquery dps-查询分货作业人员信息
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
func WdkUmsOutboundSortingUserquery(clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingUserqueryAPIRequest, session string) (*wdk.WdkUmsOutboundSortingUserqueryAPIResponse, error) {
	var resp wdk.WdkUmsOutboundSortingUserqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
