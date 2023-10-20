package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminUserOrderlistAPIResponse 获取用户订单列表 API返回值
// yunos.tvpubadmin.user.orderlist
//
// 获取用户订单列表
type YunosTvpubadminUserOrderlistAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminUserOrderlistAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminUserOrderlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminUserOrderlistAPIResponseModel).Reset()
}

// YunosTvpubadminUserOrderlistAPIResponseModel is 获取用户订单列表 成功返回结果
type YunosTvpubadminUserOrderlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_orderlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminUserOrderlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminUserOrderlistAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminUserOrderlistAPIResponse)
	},
}

// GetYunosTvpubadminUserOrderlistAPIResponse 从 sync.Pool 获取 YunosTvpubadminUserOrderlistAPIResponse
func GetYunosTvpubadminUserOrderlistAPIResponse() *YunosTvpubadminUserOrderlistAPIResponse {
	return poolYunosTvpubadminUserOrderlistAPIResponse.Get().(*YunosTvpubadminUserOrderlistAPIResponse)
}

// ReleaseYunosTvpubadminUserOrderlistAPIResponse 将 YunosTvpubadminUserOrderlistAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminUserOrderlistAPIResponse(v *YunosTvpubadminUserOrderlistAPIResponse) {
	v.Reset()
	poolYunosTvpubadminUserOrderlistAPIResponse.Put(v)
}
