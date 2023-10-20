package aecreatives

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliatecategorygetAPIResponse AE流量推广类目信息获取API API返回值
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
type AliexpressaffiliatecategorygetAPIResponse struct {
	model.CommonResponse
	AliexpressaffiliatecategorygetAPIResponseModel
}

// AliexpressaffiliatecategorygetAPIResponseModel is AE流量推广类目信息获取API 成功返回结果
type AliexpressaffiliatecategorygetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseResult `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}
