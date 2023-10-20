package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkumsoutboundsortinguserquery dps-查询分货作业人员信息
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
func Wdkumsoutboundsortinguserquery(clt *core.SDKClient, req *wdk.WdkumsoutboundsortinguserqueryAPIRequest, session string) (*wdk.WdkumsoutboundsortinguserqueryAPIResponse, error) {
	var resp wdk.WdkumsoutboundsortinguserqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
