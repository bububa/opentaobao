package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsTaskCreateAPIRequest
聚石塔短信任务创建接口 API请求
taobao.jst.sms.task.create

聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。 */
type TaobaoJstSmsTaskCreateAPIRequest struct {
	model.Params
	// 创建任务的入参
	_paramCreateSmsTaskRequest *CreateSmsTaskRequest
}

// New
