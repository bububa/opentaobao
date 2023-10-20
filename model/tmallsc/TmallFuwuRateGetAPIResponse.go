package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuRateGetAPIResponse 服务商需获取到单条服务单评价信息 API返回值
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
type TmallFuwuRateGetAPIResponse struct {
	model.CommonResponse
	TmallFuwuRateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFuwuRateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFuwuRateGetAPIResponseModel).Reset()
}

// TmallFuwuRateGetAPIResponseModel is 服务商需获取到单条服务单评价信息 成功返回结果
type TmallFuwuRateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_rate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TmallFuwuRateGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFuwuRateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallFuwuRateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFuwuRateGetAPIResponse)
	},
}

// GetTmallFuwuRateGetAPIResponse 从 sync.Pool 获取 TmallFuwuRateGetAPIResponse
func GetTmallFuwuRateGetAPIResponse() *TmallFuwuRateGetAPIResponse {
	return poolTmallFuwuRateGetAPIResponse.Get().(*TmallFuwuRateGetAPIResponse)
}

// ReleaseTmallFuwuRateGetAPIResponse 将 TmallFuwuRateGetAPIResponse 保存到 sync.Pool
func ReleaseTmallFuwuRateGetAPIResponse(v *TmallFuwuRateGetAPIResponse) {
	v.Reset()
	poolTmallFuwuRateGetAPIResponse.Put(v)
}
