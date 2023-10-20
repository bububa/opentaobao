package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskupdate 更新轻任务
// taobao.qianniu.task.update
//
// 由任务执行者调用，sub_status，tag和memo至少提供一个
func Taobaoqianniutaskupdate(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskupdateAPIRequest, session string) (*qianniu.TaobaoqianniutaskupdateAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
