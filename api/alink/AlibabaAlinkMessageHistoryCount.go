package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkMessageHistoryCount 查询消息总数
// alibaba.alink.message.history.count
//
// 查询消息总数
func AlibabaAlinkMessageHistoryCount(clt *core.SDKClient, req *alink.AlibabaAlinkMessageHistoryCountAPIRequest, session string) (*alink.AlibabaAlinkMessageHistoryCountAPIResponse, error) {
	var resp alink.AlibabaAlinkMessageHistoryCountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
