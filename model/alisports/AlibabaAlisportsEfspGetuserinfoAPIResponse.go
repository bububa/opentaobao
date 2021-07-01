package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspGetuserinfoAPIResponse
获取用户详细信息 API返回值
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息 */
type AlibabaAlisportsEfspGetuserinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsEfspGetuserinfoAPIResponseModel
}

// AlibabaAlisportsEfspGetuserinfoAPIResponseModel is 获取用户详细信息 成功返回结果
type AlibabaAlisportsEfspGetuserinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_efsp_getuserinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`
}
