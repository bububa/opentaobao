package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiItemCooperationSyncAPIResponse 商品同步 API返回值
// taobao.paimai.item.cooperation.sync
//
// 商品同步
type TaobaoPaimaiItemCooperationSyncAPIResponse struct {
	model.CommonResponse
	TaobaoPaimaiItemCooperationSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPaimaiItemCooperationSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPaimaiItemCooperationSyncAPIResponseModel).Reset()
}

// TaobaoPaimaiItemCooperationSyncAPIResponseModel is 商品同步 成功返回结果
type TaobaoPaimaiItemCooperationSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_item_cooperation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功上传的商品ID列表
	Value []int64 `json:"value,omitempty" xml:"value>int64,omitempty"`
	// 结果描述
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPaimaiItemCooperationSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = m.Value[:0]
	m.ResultCode = nil
}

var poolTaobaoPaimaiItemCooperationSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPaimaiItemCooperationSyncAPIResponse)
	},
}

// GetTaobaoPaimaiItemCooperationSyncAPIResponse 从 sync.Pool 获取 TaobaoPaimaiItemCooperationSyncAPIResponse
func GetTaobaoPaimaiItemCooperationSyncAPIResponse() *TaobaoPaimaiItemCooperationSyncAPIResponse {
	return poolTaobaoPaimaiItemCooperationSyncAPIResponse.Get().(*TaobaoPaimaiItemCooperationSyncAPIResponse)
}

// ReleaseTaobaoPaimaiItemCooperationSyncAPIResponse 将 TaobaoPaimaiItemCooperationSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPaimaiItemCooperationSyncAPIResponse(v *TaobaoPaimaiItemCooperationSyncAPIResponse) {
	v.Reset()
	poolTaobaoPaimaiItemCooperationSyncAPIResponse.Put(v)
}
