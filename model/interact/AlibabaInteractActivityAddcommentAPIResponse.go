package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityAddcommentAPIResponse 微淘评论接口 API返回值
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
type AlibabaInteractActivityAddcommentAPIResponse struct {
	model.CommonResponse
	AlibabaInteractActivityAddcommentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractActivityAddcommentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractActivityAddcommentAPIResponseModel).Reset()
}

// AlibabaInteractActivityAddcommentAPIResponseModel is 微淘评论接口 成功返回结果
type AlibabaInteractActivityAddcommentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_addcomment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评论的楼层数
	Floor int64 `json:"floor,omitempty" xml:"floor,omitempty"`
	// 返回成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractActivityAddcommentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Floor = 0
	m.IsSuccess = false
}

var poolAlibabaInteractActivityAddcommentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractActivityAddcommentAPIResponse)
	},
}

// GetAlibabaInteractActivityAddcommentAPIResponse 从 sync.Pool 获取 AlibabaInteractActivityAddcommentAPIResponse
func GetAlibabaInteractActivityAddcommentAPIResponse() *AlibabaInteractActivityAddcommentAPIResponse {
	return poolAlibabaInteractActivityAddcommentAPIResponse.Get().(*AlibabaInteractActivityAddcommentAPIResponse)
}

// ReleaseAlibabaInteractActivityAddcommentAPIResponse 将 AlibabaInteractActivityAddcommentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractActivityAddcommentAPIResponse(v *AlibabaInteractActivityAddcommentAPIResponse) {
	v.Reset()
	poolAlibabaInteractActivityAddcommentAPIResponse.Put(v)
}
