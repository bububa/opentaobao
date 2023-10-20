package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskremove 轻任务删除接口
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
func Taobaoqianniutaskremove(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskremoveAPIRequest, session string) (*qianniu.TaobaoqianniutaskremoveAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
