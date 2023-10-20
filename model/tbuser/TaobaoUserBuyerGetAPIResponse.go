package tbuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouserbuyergetAPIResponse 查询买家信息API API返回值
// taobao.user.buyer.get
//
// 查询买家信息API，只能买家类应用调用。
type TaobaouserbuyergetAPIResponse struct {
	model.CommonResponse
	TaobaouserbuyergetAPIResponseModel
}

// TaobaouserbuyergetAPIResponseModel is 查询买家信息API 成功返回结果
type TaobaouserbuyergetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_buyer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户
	User *User `json:"user,omitempty" xml:"user,omitempty"`
}
