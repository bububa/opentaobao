package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeQueueUsersMarkAPIResponse 尖货交易可购买用户标记 API返回值
// taobao.opentrade.queue.users.mark
//
// 尖货交易用户标记信息回传，根据openId标记用户可购买商品
type TaobaoOpentradeQueueUsersMarkAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeQueueUsersMarkAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeQueueUsersMarkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeQueueUsersMarkAPIResponseModel).Reset()
}

// TaobaoOpentradeQueueUsersMarkAPIResponseModel is 尖货交易可购买用户标记 成功返回结果
type TaobaoOpentradeQueueUsersMarkAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_queue_users_mark_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标记成功的用户数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeQueueUsersMarkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolTaobaoOpentradeQueueUsersMarkAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeQueueUsersMarkAPIResponse)
	},
}

// GetTaobaoOpentradeQueueUsersMarkAPIResponse 从 sync.Pool 获取 TaobaoOpentradeQueueUsersMarkAPIResponse
func GetTaobaoOpentradeQueueUsersMarkAPIResponse() *TaobaoOpentradeQueueUsersMarkAPIResponse {
	return poolTaobaoOpentradeQueueUsersMarkAPIResponse.Get().(*TaobaoOpentradeQueueUsersMarkAPIResponse)
}

// ReleaseTaobaoOpentradeQueueUsersMarkAPIResponse 将 TaobaoOpentradeQueueUsersMarkAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeQueueUsersMarkAPIResponse(v *TaobaoOpentradeQueueUsersMarkAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeQueueUsersMarkAPIResponse.Put(v)
}
