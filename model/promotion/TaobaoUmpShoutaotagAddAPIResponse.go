package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpShoutaotagAddAPIResponse 手淘定向优惠打标接口 API返回值
// taobao.ump.shoutaotag.add
//
// 手淘定向优惠的优惠标签打标接口
// 给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。
type TaobaoUmpShoutaotagAddAPIResponse struct {
	model.CommonResponse
	TaobaoUmpShoutaotagAddAPIResponseModel
}

// TaobaoUmpShoutaotagAddAPIResponseModel is 手淘定向优惠打标接口 成功返回结果
type TaobaoUmpShoutaotagAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_shoutaotag_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否打标成功
	AddResult bool `json:"add_result,omitempty" xml:"add_result,omitempty"`
}
