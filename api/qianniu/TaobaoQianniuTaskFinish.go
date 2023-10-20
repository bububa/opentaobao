package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskFinish 完成轻任务
// taobao.qianniu.task.finish
//
// 由任务执行者调用
func TaobaoQianniuTaskFinish(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskFinishAPIRequest, resp *qianniu.TaobaoQianniuTaskFinishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
