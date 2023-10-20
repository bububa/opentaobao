package cainiaocntec

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse 订单状态推送 API返回值
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
type CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse struct {
	model.CommonResponse
	CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponseModel
}

// CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponseModel is 订单状态推送 成功返回结果
type CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_locallife_communitylife_syncorderstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误信息
	EMsg string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
