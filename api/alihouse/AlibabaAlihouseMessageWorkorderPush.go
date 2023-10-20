package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousemessageworkorderpush 工单消息推送
// alibaba.alihouse.message.workorder.push
//
// 工单消息推送
func Alibabaalihousemessageworkorderpush(clt *core.SDKClient, req *alihouse.AlibabaalihousemessageworkorderpushAPIRequest, session string) (*alihouse.AlibabaalihousemessageworkorderpushAPIResponse, error) {
	var resp alihouse.AlibabaalihousemessageworkorderpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
