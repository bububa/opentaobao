package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelOrderCreateResQueryAPIResponse 分销订单查询订单创建结果 API返回值
// alitrip.xhotel.channel.order.create.res.query
//
// 针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
type AlitripXhotelChannelOrderCreateResQueryAPIResponse struct {
	model.CommonResponse
	AlitripXhotelChannelOrderCreateResQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripXhotelChannelOrderCreateResQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripXhotelChannelOrderCreateResQueryAPIResponseModel).Reset()
}

// AlitripXhotelChannelOrderCreateResQueryAPIResponseModel is 分销订单查询订单创建结果 成功返回结果
type AlitripXhotelChannelOrderCreateResQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_xhotel_channel_order_create_res_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *HbsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripXhotelChannelOrderCreateResQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripXhotelChannelOrderCreateResQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripXhotelChannelOrderCreateResQueryAPIResponse)
	},
}

// GetAlitripXhotelChannelOrderCreateResQueryAPIResponse 从 sync.Pool 获取 AlitripXhotelChannelOrderCreateResQueryAPIResponse
func GetAlitripXhotelChannelOrderCreateResQueryAPIResponse() *AlitripXhotelChannelOrderCreateResQueryAPIResponse {
	return poolAlitripXhotelChannelOrderCreateResQueryAPIResponse.Get().(*AlitripXhotelChannelOrderCreateResQueryAPIResponse)
}

// ReleaseAlitripXhotelChannelOrderCreateResQueryAPIResponse 将 AlitripXhotelChannelOrderCreateResQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripXhotelChannelOrderCreateResQueryAPIResponse(v *AlitripXhotelChannelOrderCreateResQueryAPIResponse) {
	v.Reset()
	poolAlitripXhotelChannelOrderCreateResQueryAPIResponse.Put(v)
}
