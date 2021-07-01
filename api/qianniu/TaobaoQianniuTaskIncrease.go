package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

/* TaobaoQianniuTaskIncrease
增加任务接收人接口
taobao.qianniu.task.increase

根据任务元id增加任务接收人 */
func TaobaoQianniuTaskIncrease(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskIncreaseAPIRequest, session string) (*qianniu.TaobaoQianniuTaskIncreaseAPIResponse, error) {
	var resp qianniu.TaobaoQianniuTaskIncreaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
