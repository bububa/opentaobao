package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureFolderQueryAPIResponse 淘特图片空间用户文件夹查询 API返回值
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
type AlibabaTjbPictureFolderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTjbPictureFolderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTjbPictureFolderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTjbPictureFolderQueryAPIResponseModel).Reset()
}

// AlibabaTjbPictureFolderQueryAPIResponseModel is 淘特图片空间用户文件夹查询 成功返回结果
type AlibabaTjbPictureFolderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_folder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件夹列表
	FolderList []TopFolderDto `json:"folder_list,omitempty" xml:"folder_list>top_folder_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTjbPictureFolderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FolderList = m.FolderList[:0]
}

var poolAlibabaTjbPictureFolderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTjbPictureFolderQueryAPIResponse)
	},
}

// GetAlibabaTjbPictureFolderQueryAPIResponse 从 sync.Pool 获取 AlibabaTjbPictureFolderQueryAPIResponse
func GetAlibabaTjbPictureFolderQueryAPIResponse() *AlibabaTjbPictureFolderQueryAPIResponse {
	return poolAlibabaTjbPictureFolderQueryAPIResponse.Get().(*AlibabaTjbPictureFolderQueryAPIResponse)
}

// ReleaseAlibabaTjbPictureFolderQueryAPIResponse 将 AlibabaTjbPictureFolderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTjbPictureFolderQueryAPIResponse(v *AlibabaTjbPictureFolderQueryAPIResponse) {
	v.Reset()
	poolAlibabaTjbPictureFolderQueryAPIResponse.Put(v)
}
