package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveConfirmAPIResponse
体检机构对接_体检套餐预定确认 API返回值
alibaba.alihealth.examination.reserve.confirm

向体检机构确认用户购买的体检套餐信息 */
type AlibabaAlihealthExaminationReserveConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveConfirmAPIResponseModel
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
}
