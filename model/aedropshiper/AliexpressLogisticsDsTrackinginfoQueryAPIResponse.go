package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsDsTrackinginfoQueryAPIResponse 查询物流追踪信息 API返回值
// aliexpress.logistics.ds.trackinginfo.query
//
// Dropshipper查询物流追踪信息
type AliexpressLogisticsDsTrackinginfoQueryAPIResponse struct {
	model.CommonResponse
	AliexpressLogisticsDsTrackinginfoQueryAPIResponseModel
}

// AliexpressLogisticsDsTrackinginfoQueryAPIResponseModel is 查询物流追踪信息 成功返回结果
type AliexpressLogisticsDsTrackinginfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_logistics_ds_trackinginfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 追踪详细信息列表
	Details []Details `json:"details,omitempty" xml:"details>details,omitempty"`
	// 追踪网址
	OfficialWebsite string `json:"official_website,omitempty" xml:"official_website,omitempty"`
	// error description
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
