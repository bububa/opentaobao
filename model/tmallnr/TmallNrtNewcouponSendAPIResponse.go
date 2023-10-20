package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtNewcouponSendAPIResponse 券发放接口 API返回值
// tmall.nrt.newcoupon.send
//
// 券发放接口
type TmallNrtNewcouponSendAPIResponse struct {
	model.CommonResponse
	TmallNrtNewcouponSendAPIResponseModel
}

// TmallNrtNewcouponSendAPIResponseModel is 券发放接口 成功返回结果
type TmallNrtNewcouponSendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_newcoupon_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TmallNrtNewcouponSendResult `json:"result,omitempty" xml:"result,omitempty"`
}
