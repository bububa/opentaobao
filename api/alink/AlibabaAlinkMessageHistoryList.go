package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkMessageHistoryList 查询消息列表
// alibaba.alink.message.history.list
//
// 查询消息列表
func AlibabaAlinkMessageHistoryList(clt *core.SDKClient, req *alink.AlibabaAlinkMessageHistoryListAPIRequest, resp *alink.AlibabaAlinkMessageHistoryListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
