package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseReversedetail 退款详情
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
func AlibabaWdkReverseReversedetail(clt *core.SDKClient, req *wdk.AlibabaWdkReverseReversedetailAPIRequest, session string) (*wdk.AlibabaWdkReverseReversedetailAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseReversedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
