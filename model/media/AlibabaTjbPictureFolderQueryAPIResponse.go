package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatjbpicturefolderqueryAPIResponse 淘特图片空间用户文件夹查询 API返回值
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
type AlibabatjbpicturefolderqueryAPIResponse struct {
	model.CommonResponse
	AlibabatjbpicturefolderqueryAPIResponseModel
}

// AlibabatjbpicturefolderqueryAPIResponseModel is 淘特图片空间用户文件夹查询 成功返回结果
type AlibabatjbpicturefolderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_folder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件夹列表
	FolderList []TopFolderDto `json:"folder_list,omitempty" xml:"folder_list>top_folder_dto,omitempty"`
}
