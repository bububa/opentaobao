package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskRemove 轻任务删除接口
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
func TaobaoQianniuTaskRemove(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskRemoveAPIRequest, session string) (*qianniu.TaobaoQianniuTaskRemoveAPIResponse, error) {
	var resp qianniu.TaobaoQianniuTaskRemoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
