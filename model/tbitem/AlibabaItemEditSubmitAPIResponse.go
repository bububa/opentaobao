package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditSubmitAPIResponse 商品编辑提交schema信息 API返回值
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
type AlibabaItemEditSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaItemEditSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemEditSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemEditSubmitAPIResponseModel).Reset()
}

// AlibabaItemEditSubmitAPIResponseModel is 商品编辑提交schema信息 成功返回结果
type AlibabaItemEditSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_edit_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 商品所属市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemEditSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateTime = ""
	m.Market = ""
	m.ItemId = 0
}

var poolAlibabaItemEditSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemEditSubmitAPIResponse)
	},
}

// GetAlibabaItemEditSubmitAPIResponse 从 sync.Pool 获取 AlibabaItemEditSubmitAPIResponse
func GetAlibabaItemEditSubmitAPIResponse() *AlibabaItemEditSubmitAPIResponse {
	return poolAlibabaItemEditSubmitAPIResponse.Get().(*AlibabaItemEditSubmitAPIResponse)
}

// ReleaseAlibabaItemEditSubmitAPIResponse 将 AlibabaItemEditSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemEditSubmitAPIResponse(v *AlibabaItemEditSubmitAPIResponse) {
	v.Reset()
	poolAlibabaItemEditSubmitAPIResponse.Put(v)
}
