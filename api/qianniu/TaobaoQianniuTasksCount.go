package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskscount 任务查询条数接口
// taobao.qianniu.tasks.count
//
// 任务查询条数接口
func Taobaoqianniutaskscount(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskscountAPIRequest, session string) (*qianniu.TaobaoqianniutaskscountAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskscountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
