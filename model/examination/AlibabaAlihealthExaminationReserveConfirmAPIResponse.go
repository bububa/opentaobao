package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreserveconfirmAPIResponse 体检机构对接_体检套餐预定确认 API返回值
// alibaba.alihealth.examination.reserve.confirm
//
// 向体检机构确认用户购买的体检套餐信息
type AlibabaalihealthexaminationreserveconfirmAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreserveconfirmAPIResponseModel
}

// AlibabaalihealthexaminationreserveconfirmAPIResponseModel is 体检机构对接_体检套餐预定确认 成功返回结果
type AlibabaalihealthexaminationreserveconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 体检机构预约唯一标识码
	UniqReserveCode string `json:"uniq_reserve_code,omitempty" xml:"uniq_reserve_code,omitempty"`
	// 返回结果编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 预约电子凭证码值，若返回凭证码值，会展示在订单详情。用户到店/上门后，可以通过该码值来验证订单，不同预约码值不能重复。长度不超过64位
	VoucherCode string `json:"voucher_code,omitempty" xml:"voucher_code,omitempty"`
}
