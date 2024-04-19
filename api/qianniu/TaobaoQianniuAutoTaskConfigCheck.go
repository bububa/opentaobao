package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuAutoTaskConfigCheck 自动化任务设置校验
// taobao.qianniu.auto.task.config.check
//
// 校验自动化任务配置
func TaobaoQianniuAutoTaskConfigCheck(clt *core.SDKClient, req *qianniu.TaobaoQianniuAutoTaskConfigCheckAPIRequest, resp *qianniu.TaobaoQianniuAutoTaskConfigCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
