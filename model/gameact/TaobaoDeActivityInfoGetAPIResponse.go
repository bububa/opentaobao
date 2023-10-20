package gameact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityInfoGetAPIResponse 获取活动信息 API返回值
// taobao.de.activity.info.get
//
// 根据appKey和活动id获取活动
type TaobaoDeActivityInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoDeActivityInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeActivityInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeActivityInfoGetAPIResponseModel).Reset()
}

// TaobaoDeActivityInfoGetAPIResponseModel is 获取活动信息 成功返回结果
type TaobaoDeActivityInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构
	Activities []ActivityVo `json:"activities,omitempty" xml:"activities>activity_vo,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeActivityInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Activities = m.Activities[:0]
}

var poolTaobaoDeActivityInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeActivityInfoGetAPIResponse)
	},
}

// GetTaobaoDeActivityInfoGetAPIResponse 从 sync.Pool 获取 TaobaoDeActivityInfoGetAPIResponse
func GetTaobaoDeActivityInfoGetAPIResponse() *TaobaoDeActivityInfoGetAPIResponse {
	return poolTaobaoDeActivityInfoGetAPIResponse.Get().(*TaobaoDeActivityInfoGetAPIResponse)
}

// ReleaseTaobaoDeActivityInfoGetAPIResponse 将 TaobaoDeActivityInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeActivityInfoGetAPIResponse(v *TaobaoDeActivityInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoDeActivityInfoGetAPIResponse.Put(v)
}
