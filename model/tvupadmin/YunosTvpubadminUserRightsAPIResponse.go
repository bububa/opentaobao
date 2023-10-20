package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminuserrightsAPIResponse 获取用户权益 API返回值
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
type YunostvpubadminuserrightsAPIResponse struct {
	model.CommonResponse
	YunostvpubadminuserrightsAPIResponseModel
}

// YunostvpubadminuserrightsAPIResponseModel is 获取用户权益 成功返回结果
type YunostvpubadminuserrightsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_rights_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}
