package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMessageWorkorderPush 工单消息推送
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
func AlibabaAlihouseMessageWorkorderPush(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMessageWorkorderPushAPIRequest, session string) (*alihouse.AlibabaAlihouseMessageWorkorderPushAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseMessageWorkorderPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
