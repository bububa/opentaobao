package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemUpdateAPIResponse 物流宝商品修改 API返回值
// taobao.wlb.item.update
//
// 修改物流宝商品信息
type TaobaoWlbItemUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbItemUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbItemUpdateAPIResponseModel).Reset()
}

// TaobaoWlbItemUpdateAPIResponseModel is 物流宝商品修改 成功返回结果
type TaobaoWlbItemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改时间
	GmtModified bool `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbItemUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GmtModified = false
}

var poolTaobaoWlbItemUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbItemUpdateAPIResponse)
	},
}

// GetTaobaoWlbItemUpdateAPIResponse 从 sync.Pool 获取 TaobaoWlbItemUpdateAPIResponse
func GetTaobaoWlbItemUpdateAPIResponse() *TaobaoWlbItemUpdateAPIResponse {
	return poolTaobaoWlbItemUpdateAPIResponse.Get().(*TaobaoWlbItemUpdateAPIResponse)
}

// ReleaseTaobaoWlbItemUpdateAPIResponse 将 TaobaoWlbItemUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbItemUpdateAPIResponse(v *TaobaoWlbItemUpdateAPIResponse) {
	v.Reset()
	poolTaobaoWlbItemUpdateAPIResponse.Put(v)
}
