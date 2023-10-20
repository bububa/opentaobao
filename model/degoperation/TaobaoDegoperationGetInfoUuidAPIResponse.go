package degoperation

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationGetInfoUuidAPIResponse 根据uuid用户抽奖次数限制 API返回值
// taobao.degoperation.get.info.uuid
//
// 根据uuid用户抽奖次数限制
type TaobaoDegoperationGetInfoUuidAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationGetInfoUuidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegoperationGetInfoUuidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegoperationGetInfoUuidAPIResponseModel).Reset()
}

// TaobaoDegoperationGetInfoUuidAPIResponseModel is 根据uuid用户抽奖次数限制 成功返回结果
type TaobaoDegoperationGetInfoUuidAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_get_info_uuid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegoperationGetInfoUuidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDegoperationGetInfoUuidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegoperationGetInfoUuidAPIResponse)
	},
}

// GetTaobaoDegoperationGetInfoUuidAPIResponse 从 sync.Pool 获取 TaobaoDegoperationGetInfoUuidAPIResponse
func GetTaobaoDegoperationGetInfoUuidAPIResponse() *TaobaoDegoperationGetInfoUuidAPIResponse {
	return poolTaobaoDegoperationGetInfoUuidAPIResponse.Get().(*TaobaoDegoperationGetInfoUuidAPIResponse)
}

// ReleaseTaobaoDegoperationGetInfoUuidAPIResponse 将 TaobaoDegoperationGetInfoUuidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegoperationGetInfoUuidAPIResponse(v *TaobaoDegoperationGetInfoUuidAPIResponse) {
	v.Reset()
	poolTaobaoDegoperationGetInfoUuidAPIResponse.Put(v)
}
