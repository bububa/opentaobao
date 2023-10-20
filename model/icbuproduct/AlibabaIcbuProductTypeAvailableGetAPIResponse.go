package icbuproduct

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproducttypeavailablegetAPIResponse 商家发品类型查询 API返回值
// alibaba.icbu.product.type.available.get
//
// 查询商家发品权限
type AlibabaicbuproducttypeavailablegetAPIResponse struct {
	model.CommonResponse
	AlibabaicbuproducttypeavailablegetAPIResponseModel
}

// AlibabaicbuproducttypeavailablegetAPIResponseModel is 商家发品类型查询 成功返回结果
type AlibabaicbuproducttypeavailablegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_type_available_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误追踪码，请务必打印在日志中，后续排查问题请提交此错误追踪码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的错误码，数组形式的字符串，用;分割
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回结果数据
	Data *ProductSupportTypeDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
