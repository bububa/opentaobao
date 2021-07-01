package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeRefuseVtwoAPIResponse
代理商拒绝改签v2--增加鉴权校验 API返回值
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务 */
type TaobaoTrainAgentChangeRefuseVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentChangeRefuseVtwoAPIResponseModel
}

// TaobaoTrainAgentChangeRefuseVtwoAPIResponseModel is 代理商拒绝改签v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentChangeRefuseVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_change_refuse_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
