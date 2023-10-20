package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseTimeslice 逆向取货时间片查询
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
func AlibabaWdkReverseTimeslice(clt *core.SDKClient, req *wdk.AlibabaWdkReverseTimesliceAPIRequest, session string) (*wdk.AlibabaWdkReverseTimesliceAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseTimesliceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
