package dt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaocmedicalfacedetectioncallbackAPIResponse 魔镜测肤结果数据回调接口 API返回值
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
type TaobaocmedicalfacedetectioncallbackAPIResponse struct {
	model.CommonResponse
	TaobaocmedicalfacedetectioncallbackAPIResponseModel
}

// TaobaocmedicalfacedetectioncallbackAPIResponseModel is 魔镜测肤结果数据回调接口 成功返回结果
type TaobaocmedicalfacedetectioncallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"cmedical_face_detection_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaocmedicalfacedetectioncallbackResult `json:"result,omitempty" xml:"result,omitempty"`
}
