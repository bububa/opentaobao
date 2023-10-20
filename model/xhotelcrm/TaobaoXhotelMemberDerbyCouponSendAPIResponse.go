package xhotelcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberDerbyCouponSendAPIResponse 发券 API返回值
// taobao.xhotel.member.derby.coupon.send
//
// 发券
type TaobaoXhotelMemberDerbyCouponSendAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMemberDerbyCouponSendAPIResponseModel
}

// TaobaoXhotelMemberDerbyCouponSendAPIResponseModel is 发券 成功返回结果
type TaobaoXhotelMemberDerbyCouponSendAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_member_derby_coupon_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *MsdResult `json:"result,omitempty" xml:"result,omitempty"`
}
