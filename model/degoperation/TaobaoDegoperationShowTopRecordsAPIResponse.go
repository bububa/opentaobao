package degoperation

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationShowTopRecordsAPIResponse 活动中奖记录 API返回值
// taobao.degoperation.show.top.records
//
// 活动中奖记录
type TaobaoDegoperationShowTopRecordsAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationShowTopRecordsAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegoperationShowTopRecordsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegoperationShowTopRecordsAPIResponseModel).Reset()
}

// TaobaoDegoperationShowTopRecordsAPIResponseModel is 活动中奖记录 成功返回结果
type TaobaoDegoperationShowTopRecordsAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_show_top_records_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegoperationShowTopRecordsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDegoperationShowTopRecordsAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegoperationShowTopRecordsAPIResponse)
	},
}

// GetTaobaoDegoperationShowTopRecordsAPIResponse 从 sync.Pool 获取 TaobaoDegoperationShowTopRecordsAPIResponse
func GetTaobaoDegoperationShowTopRecordsAPIResponse() *TaobaoDegoperationShowTopRecordsAPIResponse {
	return poolTaobaoDegoperationShowTopRecordsAPIResponse.Get().(*TaobaoDegoperationShowTopRecordsAPIResponse)
}

// ReleaseTaobaoDegoperationShowTopRecordsAPIResponse 将 TaobaoDegoperationShowTopRecordsAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegoperationShowTopRecordsAPIResponse(v *TaobaoDegoperationShowTopRecordsAPIResponse) {
	v.Reset()
	poolTaobaoDegoperationShowTopRecordsAPIResponse.Put(v)
}
