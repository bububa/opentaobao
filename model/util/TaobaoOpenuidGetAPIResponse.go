package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenuidgetAPIResponse 获取授权账号对应的OpenUid API返回值
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
type TaobaoopenuidgetAPIResponse struct {
	model.CommonResponse
	TaobaoopenuidgetAPIResponseModel
}

// TaobaoopenuidgetAPIResponseModel is 获取授权账号对应的OpenUid 成功返回结果
type TaobaoopenuidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"openuid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// OpenUID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}
