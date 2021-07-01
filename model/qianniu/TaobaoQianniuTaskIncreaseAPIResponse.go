package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskIncreaseAPIResponse
增加任务接收人接口 API返回值
taobao.qianniu.task.increase

根据任务元id增加任务接收人 */
type TaobaoQianniuTaskIncreaseAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskIncreaseAPIResponseModel
}

// TaobaoQianniuTaskIncreaseAPIResponseModel is 增加任务接收人接口 成功返回结果
type TaobaoQianniuTaskIncreaseAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_task_increase_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否添加成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
