package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminUserRightsAPIResponse 获取用户权益 API返回值
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
type YunosTvpubadminUserRightsAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminUserRightsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminUserRightsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminUserRightsAPIResponseModel).Reset()
}

// YunosTvpubadminUserRightsAPIResponseModel is 获取用户权益 成功返回结果
type YunosTvpubadminUserRightsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_rights_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminUserRightsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminUserRightsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminUserRightsAPIResponse)
	},
}

// GetYunosTvpubadminUserRightsAPIResponse 从 sync.Pool 获取 YunosTvpubadminUserRightsAPIResponse
func GetYunosTvpubadminUserRightsAPIResponse() *YunosTvpubadminUserRightsAPIResponse {
	return poolYunosTvpubadminUserRightsAPIResponse.Get().(*YunosTvpubadminUserRightsAPIResponse)
}

// ReleaseYunosTvpubadminUserRightsAPIResponse 将 YunosTvpubadminUserRightsAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminUserRightsAPIResponse(v *YunosTvpubadminUserRightsAPIResponse) {
	v.Reset()
	poolYunosTvpubadminUserRightsAPIResponse.Put(v)
}
