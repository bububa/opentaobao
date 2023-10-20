package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaAlinkMessageHistoryAction 操作历史消息
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
func AlibabaAlinkMessageHistoryAction(clt *core.SDKClient, req *logistic.AlibabaAlinkMessageHistoryActionAPIRequest, resp *logistic.AlibabaAlinkMessageHistoryActionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
