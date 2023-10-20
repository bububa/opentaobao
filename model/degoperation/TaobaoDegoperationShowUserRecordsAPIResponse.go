package degoperation

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationShowUserRecordsAPIResponse 用户中奖记录 API返回值
// taobao.degoperation.show.user.records
//
// 用户中奖记录
type TaobaoDegoperationShowUserRecordsAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationShowUserRecordsAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegoperationShowUserRecordsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegoperationShowUserRecordsAPIResponseModel).Reset()
}

// TaobaoDegoperationShowUserRecordsAPIResponseModel is 用户中奖记录 成功返回结果
type TaobaoDegoperationShowUserRecordsAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_show_user_records_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegoperationShowUserRecordsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDegoperationShowUserRecordsAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegoperationShowUserRecordsAPIResponse)
	},
}

// GetTaobaoDegoperationShowUserRecordsAPIResponse 从 sync.Pool 获取 TaobaoDegoperationShowUserRecordsAPIResponse
func GetTaobaoDegoperationShowUserRecordsAPIResponse() *TaobaoDegoperationShowUserRecordsAPIResponse {
	return poolTaobaoDegoperationShowUserRecordsAPIResponse.Get().(*TaobaoDegoperationShowUserRecordsAPIResponse)
}

// ReleaseTaobaoDegoperationShowUserRecordsAPIResponse 将 TaobaoDegoperationShowUserRecordsAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegoperationShowUserRecordsAPIResponse(v *TaobaoDegoperationShowUserRecordsAPIResponse) {
	v.Reset()
	poolTaobaoDegoperationShowUserRecordsAPIResponse.Put(v)
}
