package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrrxprescriptiongetAPIResponse 搜索处方详情 API返回值
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
type AlibabaalihealthnrrxprescriptiongetAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthnrrxprescriptiongetAPIResponseModel
}

// AlibabaalihealthnrrxprescriptiongetAPIResponseModel is 搜索处方详情 成功返回结果
type AlibabaalihealthnrrxprescriptiongetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_rx_prescription_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据集
	DataList []PrescriptionSearchResultDto `json:"data_list,omitempty" xml:"data_list>prescription_search_result_dto,omitempty"`
}
