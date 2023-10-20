package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOutInventoryChangeNotifyAPIResponse 外部库存变化通知（企业物流用户使用） API返回值
// taobao.wlb.out.inventory.change.notify
//
// 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOutInventoryChangeNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOutInventoryChangeNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOutInventoryChangeNotifyAPIResponseModel).Reset()
}

// TaobaoWlbOutInventoryChangeNotifyAPIResponseModel is 外部库存变化通知（企业物流用户使用） 成功返回结果
type TaobaoWlbOutInventoryChangeNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_out_inventory_change_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存变化通知成功时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOutInventoryChangeNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GmtModified = ""
}

var poolTaobaoWlbOutInventoryChangeNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOutInventoryChangeNotifyAPIResponse)
	},
}

// GetTaobaoWlbOutInventoryChangeNotifyAPIResponse 从 sync.Pool 获取 TaobaoWlbOutInventoryChangeNotifyAPIResponse
func GetTaobaoWlbOutInventoryChangeNotifyAPIResponse() *TaobaoWlbOutInventoryChangeNotifyAPIResponse {
	return poolTaobaoWlbOutInventoryChangeNotifyAPIResponse.Get().(*TaobaoWlbOutInventoryChangeNotifyAPIResponse)
}

// ReleaseTaobaoWlbOutInventoryChangeNotifyAPIResponse 将 TaobaoWlbOutInventoryChangeNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOutInventoryChangeNotifyAPIResponse(v *TaobaoWlbOutInventoryChangeNotifyAPIResponse) {
	v.Reset()
	poolTaobaoWlbOutInventoryChangeNotifyAPIResponse.Put(v)
}
