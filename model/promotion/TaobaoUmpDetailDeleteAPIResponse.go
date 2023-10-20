package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailDeleteAPIResponse 删除活动详情 API返回值
// taobao.ump.detail.delete
//
// 删除活动详情
type TaobaoUmpDetailDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoUmpDetailDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpDetailDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpDetailDeleteAPIResponseModel).Reset()
}

// TaobaoUmpDetailDeleteAPIResponseModel is 删除活动详情 成功返回结果
type TaobaoUmpDetailDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpDetailDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoUmpDetailDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpDetailDeleteAPIResponse)
	},
}

// GetTaobaoUmpDetailDeleteAPIResponse 从 sync.Pool 获取 TaobaoUmpDetailDeleteAPIResponse
func GetTaobaoUmpDetailDeleteAPIResponse() *TaobaoUmpDetailDeleteAPIResponse {
	return poolTaobaoUmpDetailDeleteAPIResponse.Get().(*TaobaoUmpDetailDeleteAPIResponse)
}

// ReleaseTaobaoUmpDetailDeleteAPIResponse 将 TaobaoUmpDetailDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpDetailDeleteAPIResponse(v *TaobaoUmpDetailDeleteAPIResponse) {
	v.Reset()
	poolTaobaoUmpDetailDeleteAPIResponse.Put(v)
}
