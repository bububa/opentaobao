package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspUsercancelorderAPIResponse
用户取消订单 API返回值
alibaba.alisports.efsp.usercancelorder

用户取消订单 */
type AlibabaAlisportsEfspUsercancelorderAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsEfspUsercancelorderAPIResponseModel
}

// AlibabaAlisportsEfspUsercancelorderAPIResponseModel is 用户取消订单 成功返回结果
type AlibabaAlisportsEfspUsercancelorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_efsp_usercancelorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
