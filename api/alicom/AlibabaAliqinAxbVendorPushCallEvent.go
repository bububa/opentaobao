package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinaxbvendorpushcallevent 呼叫事件推送
// alibaba.aliqin.axb.vendor.push.call.event
//
// 呼叫事件推送
// 响铃时间、摘机事件
func Alibabaaliqinaxbvendorpushcallevent(clt *core.SDKClient, req *alicom.AlibabaaliqinaxbvendorpushcalleventAPIRequest, session string) (*alicom.AlibabaaliqinaxbvendorpushcalleventAPIResponse, error) {
	var resp alicom.AlibabaaliqinaxbvendorpushcalleventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
