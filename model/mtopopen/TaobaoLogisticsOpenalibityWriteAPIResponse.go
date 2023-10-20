package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsopenalibitywriteAPIResponse 为快递公司提供的物流信息通用写入接口 API返回值
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
type TaobaologisticsopenalibitywriteAPIResponse struct {
	model.CommonResponse
	TaobaologisticsopenalibitywriteAPIResponseModel
}

// TaobaologisticsopenalibitywriteAPIResponseModel is 为快递公司提供的物流信息通用写入接口 成功返回结果
type TaobaologisticsopenalibitywriteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_openalibity_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用失败信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 接口调用失败码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 信息写入返回结果对象
	WriteResponse *GeneralLogisticsDataWriteResponse `json:"write_response,omitempty" xml:"write_response,omitempty"`
	// 接口调用是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
