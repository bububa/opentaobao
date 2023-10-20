package beehive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBeehiveItemCpsUrlAPIResponse 分佣链接生成接口 API返回值
// taobao.beehive.item.cps.url
//
// 传入包括itemId,accountId,bizType在内的参数，对应参数返回分佣链接
type TaobaoBeehiveItemCpsUrlAPIResponse struct {
	model.CommonResponse
	TaobaoBeehiveItemCpsUrlAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBeehiveItemCpsUrlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBeehiveItemCpsUrlAPIResponseModel).Reset()
}

// TaobaoBeehiveItemCpsUrlAPIResponseModel is 分佣链接生成接口 成功返回结果
type TaobaoBeehiveItemCpsUrlAPIResponseModel struct {
	XMLName xml.Name `xml:"beehive_item_cps_url_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *TaobaoBeehiveItemCpsUrlResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBeehiveItemCpsUrlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBeehiveItemCpsUrlAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBeehiveItemCpsUrlAPIResponse)
	},
}

// GetTaobaoBeehiveItemCpsUrlAPIResponse 从 sync.Pool 获取 TaobaoBeehiveItemCpsUrlAPIResponse
func GetTaobaoBeehiveItemCpsUrlAPIResponse() *TaobaoBeehiveItemCpsUrlAPIResponse {
	return poolTaobaoBeehiveItemCpsUrlAPIResponse.Get().(*TaobaoBeehiveItemCpsUrlAPIResponse)
}

// ReleaseTaobaoBeehiveItemCpsUrlAPIResponse 将 TaobaoBeehiveItemCpsUrlAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBeehiveItemCpsUrlAPIResponse(v *TaobaoBeehiveItemCpsUrlAPIResponse) {
	v.Reset()
	poolTaobaoBeehiveItemCpsUrlAPIResponse.Put(v)
}
