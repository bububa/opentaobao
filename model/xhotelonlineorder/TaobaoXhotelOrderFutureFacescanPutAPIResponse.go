package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderFutureFacescanPutAPIResponse 未来酒店扫脸信息上传 API返回值
// taobao.xhotel.order.future.facescan.put
//
// 未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接
type TaobaoXhotelOrderFutureFacescanPutAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderFutureFacescanPutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderFutureFacescanPutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderFutureFacescanPutAPIResponseModel).Reset()
}

// TaobaoXhotelOrderFutureFacescanPutAPIResponseModel is 未来酒店扫脸信息上传 成功返回结果
type TaobaoXhotelOrderFutureFacescanPutAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_future_facescan_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelOrderFutureFacescanPutResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderFutureFacescanPutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelOrderFutureFacescanPutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderFutureFacescanPutAPIResponse)
	},
}

// GetTaobaoXhotelOrderFutureFacescanPutAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderFutureFacescanPutAPIResponse
func GetTaobaoXhotelOrderFutureFacescanPutAPIResponse() *TaobaoXhotelOrderFutureFacescanPutAPIResponse {
	return poolTaobaoXhotelOrderFutureFacescanPutAPIResponse.Get().(*TaobaoXhotelOrderFutureFacescanPutAPIResponse)
}

// ReleaseTaobaoXhotelOrderFutureFacescanPutAPIResponse 将 TaobaoXhotelOrderFutureFacescanPutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderFutureFacescanPutAPIResponse(v *TaobaoXhotelOrderFutureFacescanPutAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderFutureFacescanPutAPIResponse.Put(v)
}
