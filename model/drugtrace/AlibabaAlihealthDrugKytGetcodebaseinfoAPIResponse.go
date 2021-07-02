package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponse 码的药品信息查询 API返回值
// alibaba.alihealth.drug.kyt.getcodebaseinfo
//
// 提供根据码查询码基本信息接口
type AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponseModel
}

// AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponseModel is 码的药品信息查询 成功返回结果
type AlibabaAlihealthDrugKytGetcodebaseinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getcodebaseinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CodeFullInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}
