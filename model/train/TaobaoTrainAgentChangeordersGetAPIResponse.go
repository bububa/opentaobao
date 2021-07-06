package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeordersGetAPIResponse 获取待改签订单 API返回值
// taobao.train.agent.changeorders.get
//
// 代理商用来获取待改签的订单列表及数量，防止代理商掉单。
type TaobaoTrainAgentChangeordersGetAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentChangeordersGetAPIResponseModel
}

// TaobaoTrainAgentChangeordersGetAPIResponseModel is 获取待改签订单 成功返回结果
type TaobaoTrainAgentChangeordersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_changeorders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 逗号连接的多个改签单id
	ApplyIds string `json:"apply_ids,omitempty" xml:"apply_ids,omitempty"`
	// 待处理订单总数量
	ApplyCount int64 `json:"apply_count,omitempty" xml:"apply_count,omitempty"`
}
