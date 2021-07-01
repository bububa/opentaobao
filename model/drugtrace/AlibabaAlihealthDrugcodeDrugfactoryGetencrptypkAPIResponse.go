package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse
获取加密公钥 API返回值
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥 */
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponseModel
}

// AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponseModel is 获取加密公钥 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_getencrptypk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 公钥证书
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
