package blackvip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoblackvipuserinfogetAPIResponse 88VIP用户信息查询 API返回值
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
type TaobaoblackvipuserinfogetAPIResponse struct {
	model.CommonResponse
	TaobaoblackvipuserinfogetAPIResponseModel
}

// TaobaoblackvipuserinfogetAPIResponseModel is 88VIP用户信息查询 成功返回结果
type TaobaoblackvipuserinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"blackvip_userinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果支持对象
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
