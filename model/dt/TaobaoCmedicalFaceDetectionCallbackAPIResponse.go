package dt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCmedicalFaceDetectionCallbackAPIResponse 魔镜测肤结果数据回调接口 API返回值
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
type TaobaoCmedicalFaceDetectionCallbackAPIResponse struct {
	model.CommonResponse
	TaobaoCmedicalFaceDetectionCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCmedicalFaceDetectionCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCmedicalFaceDetectionCallbackAPIResponseModel).Reset()
}

// TaobaoCmedicalFaceDetectionCallbackAPIResponseModel is 魔镜测肤结果数据回调接口 成功返回结果
type TaobaoCmedicalFaceDetectionCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"cmedical_face_detection_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoCmedicalFaceDetectionCallbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCmedicalFaceDetectionCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCmedicalFaceDetectionCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCmedicalFaceDetectionCallbackAPIResponse)
	},
}

// GetTaobaoCmedicalFaceDetectionCallbackAPIResponse 从 sync.Pool 获取 TaobaoCmedicalFaceDetectionCallbackAPIResponse
func GetTaobaoCmedicalFaceDetectionCallbackAPIResponse() *TaobaoCmedicalFaceDetectionCallbackAPIResponse {
	return poolTaobaoCmedicalFaceDetectionCallbackAPIResponse.Get().(*TaobaoCmedicalFaceDetectionCallbackAPIResponse)
}

// ReleaseTaobaoCmedicalFaceDetectionCallbackAPIResponse 将 TaobaoCmedicalFaceDetectionCallbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCmedicalFaceDetectionCallbackAPIResponse(v *TaobaoCmedicalFaceDetectionCallbackAPIResponse) {
	v.Reset()
	poolTaobaoCmedicalFaceDetectionCallbackAPIResponse.Put(v)
}
