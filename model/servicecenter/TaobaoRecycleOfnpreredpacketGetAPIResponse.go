package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnpreredpacketGetAPIResponse 服务商查询前置补贴红包的最新数据 API返回值
// taobao.recycle.ofnpreredpacket.get
//
// 服务商查询前置补贴红包的最新数据
type TaobaoRecycleOfnpreredpacketGetAPIResponse struct {
	model.CommonResponse
	TaobaoRecycleOfnpreredpacketGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecycleOfnpreredpacketGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecycleOfnpreredpacketGetAPIResponseModel).Reset()
}

// TaobaoRecycleOfnpreredpacketGetAPIResponseModel is 服务商查询前置补贴红包的最新数据 成功返回结果
type TaobaoRecycleOfnpreredpacketGetAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_ofnpreredpacket_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 前置补贴红包
	Data *OfnPreRedPacketDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecycleOfnpreredpacketGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecycleOfnpreredpacketGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecycleOfnpreredpacketGetAPIResponse)
	},
}

// GetTaobaoRecycleOfnpreredpacketGetAPIResponse 从 sync.Pool 获取 TaobaoRecycleOfnpreredpacketGetAPIResponse
func GetTaobaoRecycleOfnpreredpacketGetAPIResponse() *TaobaoRecycleOfnpreredpacketGetAPIResponse {
	return poolTaobaoRecycleOfnpreredpacketGetAPIResponse.Get().(*TaobaoRecycleOfnpreredpacketGetAPIResponse)
}

// ReleaseTaobaoRecycleOfnpreredpacketGetAPIResponse 将 TaobaoRecycleOfnpreredpacketGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecycleOfnpreredpacketGetAPIResponse(v *TaobaoRecycleOfnpreredpacketGetAPIResponse) {
	v.Reset()
	poolTaobaoRecycleOfnpreredpacketGetAPIResponse.Put(v)
}
