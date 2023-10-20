package wenyuvideo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuWenyuvideoSeetaGetAPIResponse 只看TA API返回值
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
type YoukuWenyuvideoSeetaGetAPIResponse struct {
	model.CommonResponse
	YoukuWenyuvideoSeetaGetAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuWenyuvideoSeetaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuWenyuvideoSeetaGetAPIResponseModel).Reset()
}

// YoukuWenyuvideoSeetaGetAPIResponseModel is 只看TA 成功返回结果
type YoukuWenyuvideoSeetaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_wenyuvideo_seeta_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *YoukuWenyuvideoSeetaGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YoukuWenyuvideoSeetaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYoukuWenyuvideoSeetaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuWenyuvideoSeetaGetAPIResponse)
	},
}

// GetYoukuWenyuvideoSeetaGetAPIResponse 从 sync.Pool 获取 YoukuWenyuvideoSeetaGetAPIResponse
func GetYoukuWenyuvideoSeetaGetAPIResponse() *YoukuWenyuvideoSeetaGetAPIResponse {
	return poolYoukuWenyuvideoSeetaGetAPIResponse.Get().(*YoukuWenyuvideoSeetaGetAPIResponse)
}

// ReleaseYoukuWenyuvideoSeetaGetAPIResponse 将 YoukuWenyuvideoSeetaGetAPIResponse 保存到 sync.Pool
func ReleaseYoukuWenyuvideoSeetaGetAPIResponse(v *YoukuWenyuvideoSeetaGetAPIResponse) {
	v.Reset()
	poolYoukuWenyuvideoSeetaGetAPIResponse.Put(v)
}
