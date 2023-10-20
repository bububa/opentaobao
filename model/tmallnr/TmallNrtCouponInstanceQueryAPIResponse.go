package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcouponinstancequeryAPIResponse 查询用户券实例 API返回值
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
type TmallnrtcouponinstancequeryAPIResponse struct {
	model.CommonResponse
	TmallnrtcouponinstancequeryAPIResponseModel
}

// TmallnrtcouponinstancequeryAPIResponseModel is 查询用户券实例 成功返回结果
type TmallnrtcouponinstancequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_instance_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallnrtcouponinstancequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
