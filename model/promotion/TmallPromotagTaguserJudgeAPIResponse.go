package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTaguserJudgeAPIResponse 用户标签判断接口 API返回值
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
type TmallPromotagTaguserJudgeAPIResponse struct {
	model.CommonResponse
	TmallPromotagTaguserJudgeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPromotagTaguserJudgeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPromotagTaguserJudgeAPIResponseModel).Reset()
}

// TmallPromotagTaguserJudgeAPIResponseModel is 用户标签判断接口 成功返回结果
type TmallPromotagTaguserJudgeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_promotag_taguser_judge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 用户是否有标签
	HasTag bool `json:"has_tag,omitempty" xml:"has_tag,omitempty"`
}

// Reset 清空结构体
func (m *TmallPromotagTaguserJudgeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
	m.HasTag = false
}

var poolTmallPromotagTaguserJudgeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPromotagTaguserJudgeAPIResponse)
	},
}

// GetTmallPromotagTaguserJudgeAPIResponse 从 sync.Pool 获取 TmallPromotagTaguserJudgeAPIResponse
func GetTmallPromotagTaguserJudgeAPIResponse() *TmallPromotagTaguserJudgeAPIResponse {
	return poolTmallPromotagTaguserJudgeAPIResponse.Get().(*TmallPromotagTaguserJudgeAPIResponse)
}

// ReleaseTmallPromotagTaguserJudgeAPIResponse 将 TmallPromotagTaguserJudgeAPIResponse 保存到 sync.Pool
func ReleaseTmallPromotagTaguserJudgeAPIResponse(v *TmallPromotagTaguserJudgeAPIResponse) {
	v.Reset()
	poolTmallPromotagTaguserJudgeAPIResponse.Put(v)
}
