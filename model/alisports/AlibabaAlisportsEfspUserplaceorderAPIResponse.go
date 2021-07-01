package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspUserplaceorderAPIResponse
用户完成支付同步订单 API返回值
alibaba.alisports.efsp.userplaceorder

用户完成支付同步订单 */
type AlibabaAlisportsEfspUserplaceorderAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsEfspUserplaceorderAPIResponseModel
}

// AlibabaAlisportsEfspUserplaceorderAPIResponseModel is 用户完成支付同步订单 成功返回结果
type AlibabaAlisportsEfspUserplaceorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_efsp_userplaceorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
