package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoeticketmerchanttbmagetAPIResponse 码商查询淘宝码接口 API返回值
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
type TaobaoeticketmerchanttbmagetAPIResponse struct {
	model.CommonResponse
	TaobaoeticketmerchanttbmagetAPIResponseModel
}

// TaobaoeticketmerchanttbmagetAPIResponseModel is 码商查询淘宝码接口 成功返回结果
type TaobaoeticketmerchanttbmagetAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_tbma_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// subCode
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// subMsg
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// respBody
	RespBody *QueryTbMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}
