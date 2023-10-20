package tvupadmin

import (
	"encoding/xml"

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

// YunosTvpubadminUserOrderlistAPIResponseModel is 获取用户订单列表 成功返回结果
type YunosTvpubadminUserOrderlistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_user_orderlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}
