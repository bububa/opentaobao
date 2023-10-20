package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse 闲鱼商品直充功能移除 API返回值
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
type AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemRechargeBatchRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemRechargeBatchRemoveAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemRechargeBatchRemoveAPIResponseModel is 闲鱼商品直充功能移除 成功返回结果
type AlibabaIdleIsvItemRechargeBatchRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_recharge_batch_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *TopListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemRechargeBatchRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse)
	},
}

// GetAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse
func GetAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse() *AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse {
	return poolAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse.Get().(*AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse)
}

// ReleaseAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse 将 AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse(v *AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemRechargeBatchRemoveAPIResponse.Put(v)
}
