package tmalltrend

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemUploadAPIResponse 天猫品牌新品同步API API返回值
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
type TmallBrandItemUploadAPIResponse struct {
	model.CommonResponse
	TmallBrandItemUploadAPIResponseModel
}

// TmallBrandItemUploadAPIResponseModel is 天猫品牌新品同步API 成功返回结果
type TmallBrandItemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_brand_item_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求参数错误
	RespErrorCode int64 `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}
