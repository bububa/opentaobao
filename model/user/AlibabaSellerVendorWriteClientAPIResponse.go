package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendorwriteclientAPIResponse 客户动态回写 API返回值
// alibaba.seller.vendor.write.client
//
// 客户动态开放API接口，外部服务商回写数据
type AlibabasellervendorwriteclientAPIResponse struct {
	model.CommonResponse
	AlibabasellervendorwriteclientAPIResponseModel
}

// AlibabasellervendorwriteclientAPIResponseModel is 客户动态回写 成功返回结果
type AlibabasellervendorwriteclientAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seller_vendor_write_client_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
	// 错误代码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 数据写入结果 true为成功
	ReturnData bool `json:"return_data,omitempty" xml:"return_data,omitempty"`
	// 服务调用结果
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}
