package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopConnectorEventPublish 连接器事件发布
// taobao.top.connector.event.publish
//
// 连接器事件发布
func TaobaoTopConnectorEventPublish(clt *core.SDKClient, req *util.TaobaoTopConnectorEventPublishAPIRequest, session string) (*util.TaobaoTopConnectorEventPublishAPIResponse, error) {
	var resp util.TaobaoTopConnectorEventPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
