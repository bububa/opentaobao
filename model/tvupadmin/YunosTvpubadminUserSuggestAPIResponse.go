package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminUserSuggestAPIResponse 获取关联账户列表 API返回值
// yunos.tvpubadmin.user.suggest
//
// 获取关联账户列表
type YunosTvpubadminUserSuggestAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminUserSuggestAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminUserSuggestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminUserSuggestAPIResponseModel).Reset()
}

// YunosTvpubadminUserSuggestAPIResponseModel is 获取关联账户列表 成功返回结果
type YunosTvpubadminUserSuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	List []AccountSuggestDo `json:"list,omitempty" xml:"list>account_suggest_do,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminUserSuggestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.List = m.List[:0]
}

var poolYunosTvpubadminUserSuggestAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminUserSuggestAPIResponse)
	},
}

// GetYunosTvpubadminUserSuggestAPIResponse 从 sync.Pool 获取 YunosTvpubadminUserSuggestAPIResponse
func GetYunosTvpubadminUserSuggestAPIResponse() *YunosTvpubadminUserSuggestAPIResponse {
	return poolYunosTvpubadminUserSuggestAPIResponse.Get().(*YunosTvpubadminUserSuggestAPIResponse)
}

// ReleaseYunosTvpubadminUserSuggestAPIResponse 将 YunosTvpubadminUserSuggestAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminUserSuggestAPIResponse(v *YunosTvpubadminUserSuggestAPIResponse) {
	v.Reset()
	poolYunosTvpubadminUserSuggestAPIResponse.Put(v)
}
