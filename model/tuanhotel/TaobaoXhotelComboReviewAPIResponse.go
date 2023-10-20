package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboReviewAPIResponse 套餐审核接口 API返回值
// taobao.xhotel.combo.review
//
// 套餐审核接口
type TaobaoXhotelComboReviewAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelComboReviewAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelComboReviewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelComboReviewAPIResponseModel).Reset()
}

// TaobaoXhotelComboReviewAPIResponseModel is 套餐审核接口 成功返回结果
type TaobaoXhotelComboReviewAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_review_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 审核状态
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelComboReviewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTaobaoXhotelComboReviewAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelComboReviewAPIResponse)
	},
}

// GetTaobaoXhotelComboReviewAPIResponse 从 sync.Pool 获取 TaobaoXhotelComboReviewAPIResponse
func GetTaobaoXhotelComboReviewAPIResponse() *TaobaoXhotelComboReviewAPIResponse {
	return poolTaobaoXhotelComboReviewAPIResponse.Get().(*TaobaoXhotelComboReviewAPIResponse)
}

// ReleaseTaobaoXhotelComboReviewAPIResponse 将 TaobaoXhotelComboReviewAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelComboReviewAPIResponse(v *TaobaoXhotelComboReviewAPIResponse) {
	v.Reset()
	poolTaobaoXhotelComboReviewAPIResponse.Put(v)
}
