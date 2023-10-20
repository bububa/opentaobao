package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskIncrease 增加任务接收人接口
// taobao.qianniu.task.increase
//
// 根据任务元id增加任务接收人
func TaobaoQianniuTaskIncrease(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskIncreaseAPIRequest, resp *qianniu.TaobaoQianniuTaskIncreaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
