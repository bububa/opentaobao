package tbuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseropenuidgetbyorderAPIResponse 根据订单获取买家openuid API返回值
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
type TaobaouseropenuidgetbyorderAPIResponse struct {
	model.CommonResponse
	TaobaouseropenuidgetbyorderAPIResponseModel
}

// TaobaouseropenuidgetbyorderAPIResponseModel is 根据订单获取买家openuid 成功返回结果
type TaobaouseropenuidgetbyorderAPIResponseModel struct {
	XMLName xml.Name `xml:"user_openuid_getbyorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 买家uid对象
	OpenUids []OpenUidInfo `json:"open_uids,omitempty" xml:"open_uids>open_uid_info,omitempty"`
}
