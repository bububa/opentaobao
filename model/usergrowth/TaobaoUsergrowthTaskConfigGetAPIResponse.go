package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthTaskConfigGetAPIResponse 用户增长营销玩法配置查询 API返回值
// taobao.usergrowth.task.config.get
//
// 用户增长营销玩法配置查询
type TaobaoUsergrowthTaskConfigGetAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthTaskConfigGetAPIResponseModel
}

// TaobaoUsergrowthTaskConfigGetAPIResponseModel is 用户增长营销玩法配置查询 成功返回结果
type TaobaoUsergrowthTaskConfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_task_config_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoUsergrowthTaskConfigGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
