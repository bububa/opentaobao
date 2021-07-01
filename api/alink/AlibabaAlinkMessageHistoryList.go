package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

/* AlibabaAlinkMessageHistoryList
查询消息列表
alibaba.alink.message.history.list

查询消息列表 */
func AlibabaAlinkMessageHistoryList(clt *core.SDKClient, req *alink.AlibabaAlinkMessageHistoryListAPIRequest, session string) (*alink.AlibabaAlinkMessageHistoryListAPIResponse, error) {
	var resp alink.AlibabaAlinkMessageHistoryListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
