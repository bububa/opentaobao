package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureQnaigcUploadAPIResponse qnaigc业务线图片上传 API返回值
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
type TaobaoPictureQnaigcUploadAPIResponse struct {
	model.CommonResponse
	TaobaoPictureQnaigcUploadAPIResponseModel
}

// TaobaoPictureQnaigcUploadAPIResponseModel is qnaigc业务线图片上传 成功返回结果
type TaobaoPictureQnaigcUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_qnaigc_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传结果
	Result *TaobaoPictureQnaigcUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
