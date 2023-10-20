package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaalinkmessagehistoryaction 操作历史消息
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
func Alibabaalinkmessagehistoryaction(clt *core.SDKClient, req *logistic.AlibabaalinkmessagehistoryactionAPIRequest, session string) (*logistic.AlibabaalinkmessagehistoryactionAPIResponse, error) {
	var resp logistic.AlibabaalinkmessagehistoryactionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
