package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelQuotaUpdateAPIResponse 库存更新接口 API返回值
// taobao.xhotel.quota.update
//
// 库存更新接口
type TaobaoXhotelQuotaUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelQuotaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelQuotaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelQuotaUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelQuotaUpdateAPIResponseModel is 库存更新接口 成功返回结果
type TaobaoXhotelQuotaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_quota_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新失败补充描述消息
	WarnMessage string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
	// errorCode
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 更新失败错误信息
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelQuotaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WarnMessage = ""
	m.BizErrorCode = ""
	m.BizErrorMsg = ""
}

var poolTaobaoXhotelQuotaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelQuotaUpdateAPIResponse)
	},
}

// GetTaobaoXhotelQuotaUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelQuotaUpdateAPIResponse
func GetTaobaoXhotelQuotaUpdateAPIResponse() *TaobaoXhotelQuotaUpdateAPIResponse {
	return poolTaobaoXhotelQuotaUpdateAPIResponse.Get().(*TaobaoXhotelQuotaUpdateAPIResponse)
}

// ReleaseTaobaoXhotelQuotaUpdateAPIResponse 将 TaobaoXhotelQuotaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelQuotaUpdateAPIResponse(v *TaobaoXhotelQuotaUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelQuotaUpdateAPIResponse.Put(v)
}
