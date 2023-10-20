package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveConfirmAPIResponse 体检机构对接_体检套餐预定确认 API返回值
// alibaba.alihealth.examination.reserve.confirm
//
// 向体检机构确认用户购买的体检套餐信息
type AlibabaAlihealthExaminationReserveConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveConfirmAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveConfirmAPIResponseModel is 体检机构对接_体检套餐预定确认 成功返回结果
type AlibabaAlihealthExaminationReserveConfirmAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.UniqReserveCode = ""
	m.ResponseCode = ""
	m.VoucherCode = ""
}

var poolAlibabaAlihealthExaminationReserveConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveConfirmAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveConfirmAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveConfirmAPIResponse
func GetAlibabaAlihealthExaminationReserveConfirmAPIResponse() *AlibabaAlihealthExaminationReserveConfirmAPIResponse {
	return poolAlibabaAlihealthExaminationReserveConfirmAPIResponse.Get().(*AlibabaAlihealthExaminationReserveConfirmAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveConfirmAPIResponse 将 AlibabaAlihealthExaminationReserveConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveConfirmAPIResponse(v *AlibabaAlihealthExaminationReserveConfirmAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveConfirmAPIResponse.Put(v)
}
