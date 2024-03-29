package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmShopvipCancelAPIResponse 卖家取消店铺vip的优惠 API返回值
// taobao.crm.shopvip.cancel
//
// 此接口用于取消VIP优惠
type TaobaoCrmShopvipCancelAPIResponse struct {
	model.CommonResponse
	TaobaoCrmShopvipCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmShopvipCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmShopvipCancelAPIResponseModel).Reset()
}

// TaobaoCrmShopvipCancelAPIResponseModel is 卖家取消店铺vip的优惠 成功返回结果
type TaobaoCrmShopvipCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_shopvip_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmShopvipCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmShopvipCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmShopvipCancelAPIResponse)
	},
}

// GetTaobaoCrmShopvipCancelAPIResponse 从 sync.Pool 获取 TaobaoCrmShopvipCancelAPIResponse
func GetTaobaoCrmShopvipCancelAPIResponse() *TaobaoCrmShopvipCancelAPIResponse {
	return poolTaobaoCrmShopvipCancelAPIResponse.Get().(*TaobaoCrmShopvipCancelAPIResponse)
}

// ReleaseTaobaoCrmShopvipCancelAPIResponse 将 TaobaoCrmShopvipCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmShopvipCancelAPIResponse(v *TaobaoCrmShopvipCancelAPIResponse) {
	v.Reset()
	poolTaobaoCrmShopvipCancelAPIResponse.Put(v)
}
