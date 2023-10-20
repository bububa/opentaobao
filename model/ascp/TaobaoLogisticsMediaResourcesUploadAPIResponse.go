package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsmediaresourcesuploadAPIResponse 图片与视频上传 API返回值
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
type TaobaologisticsmediaresourcesuploadAPIResponse struct {
	model.CommonResponse
	TaobaologisticsmediaresourcesuploadAPIResponseModel
}

// TaobaologisticsmediaresourcesuploadAPIResponseModel is 图片与视频上传 成功返回结果
type TaobaologisticsmediaresourcesuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_media_resources_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片视频上传出参
	Result *LspTopResponse `json:"result,omitempty" xml:"result,omitempty"`
}
