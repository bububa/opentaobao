package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* AlibabaAlinkMessageHistoryAction
操作历史消息
alibaba.alink.message.history.action

阿里智能操作历史消息 */
func AlibabaAlinkMessageHistoryAction(clt *core.SDKClient, req *logistic.AlibabaAlinkMessageHistoryActionAPIRequest, session string) (*logistic.AlibabaAlinkMessageHistoryActionAPIResponse, error) {
	var resp logistic.AlibabaAlinkMessageHistoryActionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
