package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemRechargeEditAPIResponse 闲鱼商品直充功能编辑 API返回值
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
type AlibabaIdleIsvItemRechargeEditAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvItemRechargeEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemRechargeEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvItemRechargeEditAPIResponseModel).Reset()
}

// AlibabaIdleIsvItemRechargeEditAPIResponseModel is 闲鱼商品直充功能编辑 成功返回结果
type AlibabaIdleIsvItemRechargeEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_recharge_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaIdleIsvItemRechargeEditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvItemRechargeEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvItemRechargeEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemRechargeEditAPIResponse)
	},
}

// GetAlibabaIdleIsvItemRechargeEditAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvItemRechargeEditAPIResponse
func GetAlibabaIdleIsvItemRechargeEditAPIResponse() *AlibabaIdleIsvItemRechargeEditAPIResponse {
	return poolAlibabaIdleIsvItemRechargeEditAPIResponse.Get().(*AlibabaIdleIsvItemRechargeEditAPIResponse)
}

// ReleaseAlibabaIdleIsvItemRechargeEditAPIResponse 将 AlibabaIdleIsvItemRechargeEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvItemRechargeEditAPIResponse(v *AlibabaIdleIsvItemRechargeEditAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvItemRechargeEditAPIResponse.Put(v)
}
