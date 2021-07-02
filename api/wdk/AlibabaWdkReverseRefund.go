package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkReverseRefund 退款打款
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
func AlibabaWdkReverseRefund(clt *core.SDKClient, req *wdk.AlibabaWdkReverseRefundAPIRequest, session string) (*wdk.AlibabaWdkReverseRefundAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
