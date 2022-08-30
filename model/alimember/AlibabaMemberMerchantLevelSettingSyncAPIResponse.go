package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberMerchantLevelSettingSyncAPIResponse 商家等级列表同步配置 API返回值
// alibaba.member.merchant.level.setting.sync
//
// 商家等级列表同步配置
type AlibabaMemberMerchantLevelSettingSyncAPIResponse struct {
	model.CommonResponse
	AlibabaMemberMerchantLevelSettingSyncAPIResponseModel
}

// AlibabaMemberMerchantLevelSettingSyncAPIResponseModel is 商家等级列表同步配置 成功返回结果
type AlibabaMemberMerchantLevelSettingSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_merchant_level_setting_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 务端系统异错误码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
