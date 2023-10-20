package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskincrease 增加任务接收人接口
// taobao.qianniu.task.increase
//
// 根据任务元id增加任务接收人
func Taobaoqianniutaskincrease(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskincreaseAPIRequest, session string) (*qianniu.TaobaoqianniutaskincreaseAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskincreaseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
