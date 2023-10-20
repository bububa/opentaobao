package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTjbPictureUserstorageQueryAPIResponse 淘特图片空间用户容量查询 API返回值
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
type AlibabaTjbPictureUserstorageQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTjbPictureUserstorageQueryAPIResponseModel
}

// AlibabaTjbPictureUserstorageQueryAPIResponseModel is 淘特图片空间用户容量查询 成功返回结果
type AlibabaTjbPictureUserstorageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tjb_picture_userstorage_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户容量信息
	UserStorage *TopUserStorageDto `json:"user_storage,omitempty" xml:"user_storage,omitempty"`
}
