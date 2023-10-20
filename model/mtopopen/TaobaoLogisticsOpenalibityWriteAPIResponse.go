package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOpenalibityWriteAPIResponse 为快递公司提供的物流信息通用写入接口 API返回值
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
type TaobaoLogisticsOpenalibityWriteAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOpenalibityWriteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsOpenalibityWriteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsOpenalibityWriteAPIResponseModel).Reset()
}

// TaobaoLogisticsOpenalibityWriteAPIResponseModel is 为快递公司提供的物流信息通用写入接口 成功返回结果
type TaobaoLogisticsOpenalibityWriteAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoLogisticsOpenalibityWriteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.WriteResponse = nil
	m.ResultSuccess = false
}

var poolTaobaoLogisticsOpenalibityWriteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsOpenalibityWriteAPIResponse)
	},
}

// GetTaobaoLogisticsOpenalibityWriteAPIResponse 从 sync.Pool 获取 TaobaoLogisticsOpenalibityWriteAPIResponse
func GetTaobaoLogisticsOpenalibityWriteAPIResponse() *TaobaoLogisticsOpenalibityWriteAPIResponse {
	return poolTaobaoLogisticsOpenalibityWriteAPIResponse.Get().(*TaobaoLogisticsOpenalibityWriteAPIResponse)
}

// ReleaseTaobaoLogisticsOpenalibityWriteAPIResponse 将 TaobaoLogisticsOpenalibityWriteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsOpenalibityWriteAPIResponse(v *TaobaoLogisticsOpenalibityWriteAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsOpenalibityWriteAPIResponse.Put(v)
}
