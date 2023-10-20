package usergrowth

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoUsergrowthTaskConfigGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsergrowthTaskConfigGetAPIResponseModel).Reset()
}

// TaobaoUsergrowthTaskConfigGetAPIResponseModel is 用户增长营销玩法配置查询 成功返回结果
type TaobaoUsergrowthTaskConfigGetAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_task_config_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoUsergrowthTaskConfigGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsergrowthTaskConfigGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsergrowthTaskConfigGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsergrowthTaskConfigGetAPIResponse)
	},
}

// GetTaobaoUsergrowthTaskConfigGetAPIResponse 从 sync.Pool 获取 TaobaoUsergrowthTaskConfigGetAPIResponse
func GetTaobaoUsergrowthTaskConfigGetAPIResponse() *TaobaoUsergrowthTaskConfigGetAPIResponse {
	return poolTaobaoUsergrowthTaskConfigGetAPIResponse.Get().(*TaobaoUsergrowthTaskConfigGetAPIResponse)
}

// ReleaseTaobaoUsergrowthTaskConfigGetAPIResponse 将 TaobaoUsergrowthTaskConfigGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsergrowthTaskConfigGetAPIResponse(v *TaobaoUsergrowthTaskConfigGetAPIResponse) {
	v.Reset()
	poolTaobaoUsergrowthTaskConfigGetAPIResponse.Put(v)
}
