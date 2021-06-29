package alidoc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
处方详情 API返回值 
alibaba.alihealth.rx.prescription.detail

获取处方结构化信息
*/
type AlibabaAlihealthRxPrescriptionDetailAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthRxPrescriptionDetailResponse
}

// 处方详情 成功返回结果
type AlibabaAlihealthRxPrescriptionDetailResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_rx_prescription_detail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 数据集
    DataList   []RxPrescriptionTopDTO `json:"data_list,omitempty" xml:"data_list>rx_prescription_top_dto,omitempty"`
}
