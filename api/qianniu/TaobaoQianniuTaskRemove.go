package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskRemove 轻任务删除接口
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
func TaobaoQianniuTaskRemove(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskRemoveAPIRequest, resp *qianniu.TaobaoQianniuTaskRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
