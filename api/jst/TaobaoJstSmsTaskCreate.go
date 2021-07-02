package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsTaskCreate 聚石塔短信任务创建接口
// taobao.jst.sms.task.create
//
// 聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
func TaobaoJstSmsTaskCreate(clt *core.SDKClient, req *jst.TaobaoJstSmsTaskCreateAPIRequest, session string) (*jst.TaobaoJstSmsTaskCreateAPIResponse, error) {
	var resp jst.TaobaoJstSmsTaskCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
