package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponse 药厂传输数据 API返回值
// alibaba.alihealth.drugcode.drugfactory.transferdata
//
// 药厂传输数据
type AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponseModel
}

// AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponseModel is 药厂传输数据 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryTransferdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_transferdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否验签成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
