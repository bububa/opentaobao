package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseCreatefeatch 逆向取货
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
func AlibabaWdkReverseCreatefeatch(clt *core.SDKClient, req *wdk.AlibabaWdkReverseCreatefeatchAPIRequest, session string) (*wdk.AlibabaWdkReverseCreatefeatchAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseCreatefeatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
