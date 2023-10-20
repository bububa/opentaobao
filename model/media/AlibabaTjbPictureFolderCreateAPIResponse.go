package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureFolderCreateAPIResponse 淘特图片空间创建文件夹 API返回值
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
type AlibabaTjbPictureFolderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaTjbPictureFolderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTjbPictureFolderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTjbPictureFolderCreateAPIResponseModel).Reset()
}

// AlibabaTjbPictureFolderCreateAPIResponseModel is 淘特图片空间创建文件夹 成功返回结果
type AlibabaTjbPictureFolderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_folder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建文件夹的返回数据
	Data *TopCreateFolderDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTjbPictureFolderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaTjbPictureFolderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTjbPictureFolderCreateAPIResponse)
	},
}

// GetAlibabaTjbPictureFolderCreateAPIResponse 从 sync.Pool 获取 AlibabaTjbPictureFolderCreateAPIResponse
func GetAlibabaTjbPictureFolderCreateAPIResponse() *AlibabaTjbPictureFolderCreateAPIResponse {
	return poolAlibabaTjbPictureFolderCreateAPIResponse.Get().(*AlibabaTjbPictureFolderCreateAPIResponse)
}

// ReleaseAlibabaTjbPictureFolderCreateAPIResponse 将 AlibabaTjbPictureFolderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTjbPictureFolderCreateAPIResponse(v *AlibabaTjbPictureFolderCreateAPIResponse) {
	v.Reset()
	poolAlibabaTjbPictureFolderCreateAPIResponse.Put(v)
}
