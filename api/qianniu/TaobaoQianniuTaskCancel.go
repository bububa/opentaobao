package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskcancel 取消轻任务
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
func Taobaoqianniutaskcancel(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskcancelAPIRequest, session string) (*qianniu.TaobaoqianniutaskcancelAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
