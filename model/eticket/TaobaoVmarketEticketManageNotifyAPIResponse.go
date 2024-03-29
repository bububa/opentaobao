package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketManageNotifyAPIResponse 主动发起通知接口 API返回值
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
type TaobaoVmarketEticketManageNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketManageNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketManageNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketManageNotifyAPIResponseModel).Reset()
}

// TaobaoVmarketEticketManageNotifyAPIResponseModel is 主动发起通知接口 成功返回结果
type TaobaoVmarketEticketManageNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_manage_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketManageNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketManageNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketManageNotifyAPIResponse)
	},
}

// GetTaobaoVmarketEticketManageNotifyAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketManageNotifyAPIResponse
func GetTaobaoVmarketEticketManageNotifyAPIResponse() *TaobaoVmarketEticketManageNotifyAPIResponse {
	return poolTaobaoVmarketEticketManageNotifyAPIResponse.Get().(*TaobaoVmarketEticketManageNotifyAPIResponse)
}

// ReleaseTaobaoVmarketEticketManageNotifyAPIResponse 将 TaobaoVmarketEticketManageNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketManageNotifyAPIResponse(v *TaobaoVmarketEticketManageNotifyAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketManageNotifyAPIResponse.Put(v)
}
