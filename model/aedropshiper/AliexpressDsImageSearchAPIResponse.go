package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsimagesearchAPIResponse 图片搜索 API返回值
// aliexpress.ds.image.search
//
// 图片搜索
type AliexpressdsimagesearchAPIResponse struct {
	model.CommonResponse
	AliexpressdsimagesearchAPIResponseModel
}

// AliexpressdsimagesearchAPIResponseModel is 图片搜索 成功返回结果
type AliexpressdsimagesearchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_image_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// System Error
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// result object
	Data *TrafficImageSearchResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// total record count
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
}
