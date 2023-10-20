package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopEventPublish 同步事件发布
// taobao.top.event.publish
//
// 同步事件发布
func TaobaoTopEventPublish(clt *core.SDKClient, req *util.TaobaoTopEventPublishAPIRequest, resp *util.TaobaoTopEventPublishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
