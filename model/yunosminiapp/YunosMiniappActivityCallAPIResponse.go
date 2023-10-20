package yunosminiapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosMiniappActivityCallAPIResponse 调用活动接口 API返回值
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
type YunosMiniappActivityCallAPIResponse struct {
	model.CommonResponse
	YunosMiniappActivityCallAPIResponseModel
}

// Reset 清空结构体
func (m *YunosMiniappActivityCallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosMiniappActivityCallAPIResponseModel).Reset()
}

// YunosMiniappActivityCallAPIResponseModel is 调用活动接口 成功返回结果
type YunosMiniappActivityCallAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_miniapp_activity_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *YunosMiniappActivityCallResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosMiniappActivityCallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosMiniappActivityCallAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosMiniappActivityCallAPIResponse)
	},
}

// GetYunosMiniappActivityCallAPIResponse 从 sync.Pool 获取 YunosMiniappActivityCallAPIResponse
func GetYunosMiniappActivityCallAPIResponse() *YunosMiniappActivityCallAPIResponse {
	return poolYunosMiniappActivityCallAPIResponse.Get().(*YunosMiniappActivityCallAPIResponse)
}

// ReleaseYunosMiniappActivityCallAPIResponse 将 YunosMiniappActivityCallAPIResponse 保存到 sync.Pool
func ReleaseYunosMiniappActivityCallAPIResponse(v *YunosMiniappActivityCallAPIResponse) {
	v.Reset()
	poolYunosMiniappActivityCallAPIResponse.Put(v)
}
