package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskIncreaseAPIRequest
增加任务接收人接口 API请求
taobao.qianniu.task.increase

根据任务元id增加任务接收人 */
type TaobaoQianniuTaskIncreaseAPIRequest struct {
	model.Params
	// 任务元id
	_metadataId int64
	// 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
	_tasks string
}

// New
