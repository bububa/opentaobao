package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse 同步外投任务 API返回值
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
type TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappExtDeliverySellerTaskSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappExtDeliverySellerTaskSyncAPIResponseModel).Reset()
}

// TaobaoMiniappExtDeliverySellerTaskSyncAPIResponseModel is 同步外投任务 成功返回结果
type TaobaoMiniappExtDeliverySellerTaskSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_ext_delivery_seller_task_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 任务id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappExtDeliverySellerTaskSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorType = 0
	m.Model = 0
	m.Successful = false
}

var poolTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse)
	},
}

// GetTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse
func GetTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse() *TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse {
	return poolTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse.Get().(*TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse)
}

// ReleaseTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse 将 TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse(v *TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellerTaskSyncAPIResponse.Put(v)
}
