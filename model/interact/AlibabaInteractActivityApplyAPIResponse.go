package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityApplyAPIResponse ISV报名官方活动(中心化流量) API返回值
// alibaba.interact.activity.apply
//
// 支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出
type AlibabaInteractActivityApplyAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractActivityApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractActivityApplyAPIResponseModel).Reset()
}

// AlibabaInteractActivityApplyAPIResponseModel is ISV报名官方活动(中心化流量) 成功返回结果
type AlibabaInteractActivityApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出错提示信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 服务结果对象
	Data *ActivityWriteResult `json:"data,omitempty" xml:"data,omitempty"`
	// top接口执行成功与否
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractActivityApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrMsg = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolAlibabaInteractActivityApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractActivityApplyAPIResponse)
	},
}

// GetAlibabaInteractActivityApplyAPIResponse 从 sync.Pool 获取 AlibabaInteractActivityApplyAPIResponse
func GetAlibabaInteractActivityApplyAPIResponse() *AlibabaInteractActivityApplyAPIResponse {
	return poolAlibabaInteractActivityApplyAPIResponse.Get().(*AlibabaInteractActivityApplyAPIResponse)
}

// ReleaseAlibabaInteractActivityApplyAPIResponse 将 AlibabaInteractActivityApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractActivityApplyAPIResponse(v *AlibabaInteractActivityApplyAPIResponse) {
	v.Reset()
	poolAlibabaInteractActivityApplyAPIResponse.Put(v)
}
