package xhotelcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelpotentialmemberbindAPIResponse 飞猪酒店商家会员绑定 API返回值
// taobao.xhotel.potential.member.bind
//
// 支持互通商家发起会员绑定
type TaobaoxhotelpotentialmemberbindAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelpotentialmemberbindAPIResponseModel
}

// TaobaoxhotelpotentialmemberbindAPIResponseModel is 飞猪酒店商家会员绑定 成功返回结果
type TaobaoxhotelpotentialmemberbindAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_potential_member_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
