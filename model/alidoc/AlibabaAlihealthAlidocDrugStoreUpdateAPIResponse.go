package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse
更新药店 API返回值
alibaba.alihealth.alidoc.drug.store.update

药店信息更新接口 */
type AlibabaAlihealthAlidocDrugStoreUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlidocDrugStoreUpdateAPIResponseModel
}

// AlibabaAlihealthAlidocDrugStoreUpdateAPIResponseModel is 更新药店 成功返回结果
type AlibabaAlihealthAlidocDrugStoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alidoc_drug_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrKode string `json:"err_kode,omitempty" xml:"err_kode,omitempty"`
	// success
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}
