package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthalidocdrugstoreaddAPIResponse gsk新增药店 API返回值
// alibaba.alihealth.alidoc.drug.store.add
//
// GSK上传药店信息
type AlibabaalihealthalidocdrugstoreaddAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthalidocdrugstoreaddAPIResponseModel
}

// AlibabaalihealthalidocdrugstoreaddAPIResponseModel is gsk新增药店 成功返回结果
type AlibabaalihealthalidocdrugstoreaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alidoc_drug_store_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errCode
	ErrorKode string `json:"error_kode,omitempty" xml:"error_kode,omitempty"`
	// errMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}
