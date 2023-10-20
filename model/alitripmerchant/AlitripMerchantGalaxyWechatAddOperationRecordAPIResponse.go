package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechataddoperationrecordAPIResponse 用户领取会员卡记录接口 API返回值
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
type AlitripmerchantgalaxywechataddoperationrecordAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechataddoperationrecordAPIResponseModel
}

// AlitripmerchantgalaxywechataddoperationrecordAPIResponseModel is 用户领取会员卡记录接口 成功返回结果
type AlitripmerchantgalaxywechataddoperationrecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_add_operation_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlitripmerchantgalaxywechataddoperationrecordResponse `json:"result,omitempty" xml:"result,omitempty"`
}
