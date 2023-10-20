package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsInventoryLackUploadAPIResponse 缺货通知 API返回值
// taobao.wlb.wms.inventory.lack.upload
//
// 缺货通知
type TaobaoWlbWmsInventoryLackUploadAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsInventoryLackUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsInventoryLackUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsInventoryLackUploadAPIResponseModel).Reset()
}

// TaobaoWlbWmsInventoryLackUploadAPIResponseModel is 缺货通知 成功返回结果
type TaobaoWlbWmsInventoryLackUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_inventory_lack_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 缺货回告
	Result *WlbWmsInventoryLackUploadResp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsInventoryLackUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbWmsInventoryLackUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsInventoryLackUploadAPIResponse)
	},
}

// GetTaobaoWlbWmsInventoryLackUploadAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsInventoryLackUploadAPIResponse
func GetTaobaoWlbWmsInventoryLackUploadAPIResponse() *TaobaoWlbWmsInventoryLackUploadAPIResponse {
	return poolTaobaoWlbWmsInventoryLackUploadAPIResponse.Get().(*TaobaoWlbWmsInventoryLackUploadAPIResponse)
}

// ReleaseTaobaoWlbWmsInventoryLackUploadAPIResponse 将 TaobaoWlbWmsInventoryLackUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsInventoryLackUploadAPIResponse(v *TaobaoWlbWmsInventoryLackUploadAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsInventoryLackUploadAPIResponse.Put(v)
}
