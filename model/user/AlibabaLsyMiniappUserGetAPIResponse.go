package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyMiniappUserGetAPIResponse 零售云小程序获取登录用户信息 API返回值
// alibaba.lsy.miniapp.user.get
//
// 零售云小程序，通过授权码获取登录的卖家账号信息
type AlibabaLsyMiniappUserGetAPIResponse struct {
	model.CommonResponse
	AlibabaLsyMiniappUserGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyMiniappUserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyMiniappUserGetAPIResponseModel).Reset()
}

// AlibabaLsyMiniappUserGetAPIResponseModel is 零售云小程序获取登录用户信息 成功返回结果
type AlibabaLsyMiniappUserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_miniapp_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应内容
	Result *MiniAppResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyMiniappUserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyMiniappUserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyMiniappUserGetAPIResponse)
	},
}

// GetAlibabaLsyMiniappUserGetAPIResponse 从 sync.Pool 获取 AlibabaLsyMiniappUserGetAPIResponse
func GetAlibabaLsyMiniappUserGetAPIResponse() *AlibabaLsyMiniappUserGetAPIResponse {
	return poolAlibabaLsyMiniappUserGetAPIResponse.Get().(*AlibabaLsyMiniappUserGetAPIResponse)
}

// ReleaseAlibabaLsyMiniappUserGetAPIResponse 将 AlibabaLsyMiniappUserGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyMiniappUserGetAPIResponse(v *AlibabaLsyMiniappUserGetAPIResponse) {
	v.Reset()
	poolAlibabaLsyMiniappUserGetAPIResponse.Put(v)
}
