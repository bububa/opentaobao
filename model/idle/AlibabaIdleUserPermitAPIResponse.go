package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitAPIResponse 用户appkey授权 API返回值
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
type AlibabaIdleUserPermitAPIResponse struct {
	model.CommonResponse
	AlibabaIdleUserPermitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleUserPermitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleUserPermitAPIResponseModel).Reset()
}

// AlibabaIdleUserPermitAPIResponseModel is 用户appkey授权 成功返回结果
type AlibabaIdleUserPermitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_user_permit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleUserPermitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleUserPermitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleUserPermitAPIResponse)
	},
}

// GetAlibabaIdleUserPermitAPIResponse 从 sync.Pool 获取 AlibabaIdleUserPermitAPIResponse
func GetAlibabaIdleUserPermitAPIResponse() *AlibabaIdleUserPermitAPIResponse {
	return poolAlibabaIdleUserPermitAPIResponse.Get().(*AlibabaIdleUserPermitAPIResponse)
}

// ReleaseAlibabaIdleUserPermitAPIResponse 将 AlibabaIdleUserPermitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleUserPermitAPIResponse(v *AlibabaIdleUserPermitAPIResponse) {
	v.Reset()
	poolAlibabaIdleUserPermitAPIResponse.Put(v)
}
