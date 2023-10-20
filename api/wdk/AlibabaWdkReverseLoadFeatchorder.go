package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseLoadFeatchorder 取货详情
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
func AlibabaWdkReverseLoadFeatchorder(clt *core.SDKClient, req *wdk.AlibabaWdkReverseLoadFeatchorderAPIRequest, session string) (*wdk.AlibabaWdkReverseLoadFeatchorderAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseLoadFeatchorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
