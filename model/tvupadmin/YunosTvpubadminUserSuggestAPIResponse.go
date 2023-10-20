package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminusersuggestAPIResponse 获取关联账户列表 API返回值
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
type YunostvpubadminusersuggestAPIResponse struct {
	model.CommonResponse
	YunostvpubadminusersuggestAPIResponseModel
}

// YunostvpubadminusersuggestAPIResponseModel is 获取关联账户列表 成功返回结果
type YunostvpubadminusersuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	List []AccountSuggestDo `json:"list,omitempty" xml:"list>account_suggest_do,omitempty"`
}
