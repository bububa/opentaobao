package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminCommonFileUploadAPIResponse 文件上传API API返回值
// yunos.tvpubadmin.common.file.upload
//
// 文件上传服务
type YunosTvpubadminCommonFileUploadAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminCommonFileUploadAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminCommonFileUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminCommonFileUploadAPIResponseModel).Reset()
}

// YunosTvpubadminCommonFileUploadAPIResponseModel is 文件上传API 成功返回结果
type YunosTvpubadminCommonFileUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_common_file_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传后的文件地址
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminCommonFileUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminCommonFileUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminCommonFileUploadAPIResponse)
	},
}

// GetYunosTvpubadminCommonFileUploadAPIResponse 从 sync.Pool 获取 YunosTvpubadminCommonFileUploadAPIResponse
func GetYunosTvpubadminCommonFileUploadAPIResponse() *YunosTvpubadminCommonFileUploadAPIResponse {
	return poolYunosTvpubadminCommonFileUploadAPIResponse.Get().(*YunosTvpubadminCommonFileUploadAPIResponse)
}

// ReleaseYunosTvpubadminCommonFileUploadAPIResponse 将 YunosTvpubadminCommonFileUploadAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminCommonFileUploadAPIResponse(v *YunosTvpubadminCommonFileUploadAPIResponse) {
	v.Reset()
	poolYunosTvpubadminCommonFileUploadAPIResponse.Put(v)
}
