package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMessageWorkorderPush 工单消息推送
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
func AlibabaAlihouseMessageWorkorderPush(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMessageWorkorderPushAPIRequest, resp *alihouse.AlibabaAlihouseMessageWorkorderPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
