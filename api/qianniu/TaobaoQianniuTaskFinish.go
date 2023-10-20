package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskfinish 完成轻任务
// taobao.qianniu.task.finish
//
// 由任务执行者调用
func Taobaoqianniutaskfinish(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskfinishAPIRequest, session string) (*qianniu.TaobaoqianniutaskfinishAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskfinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
