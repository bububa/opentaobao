package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemPublishSubmitAPIResponse 商品发布 API返回值
// alibaba.item.publish.submit
//
// 新商品发布，提交商品发布信息
type AlibabaItemPublishSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaItemPublishSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItemPublishSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItemPublishSubmitAPIResponseModel).Reset()
}

// AlibabaItemPublishSubmitAPIResponseModel is 商品发布 成功返回结果
type AlibabaItemPublishSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_publish_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 商品所属市场
	Market string `json:"market,omitempty" xml:"market,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItemPublishSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CreateTime = ""
	m.Market = ""
	m.ItemId = 0
}

var poolAlibabaItemPublishSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItemPublishSubmitAPIResponse)
	},
}

// GetAlibabaItemPublishSubmitAPIResponse 从 sync.Pool 获取 AlibabaItemPublishSubmitAPIResponse
func GetAlibabaItemPublishSubmitAPIResponse() *AlibabaItemPublishSubmitAPIResponse {
	return poolAlibabaItemPublishSubmitAPIResponse.Get().(*AlibabaItemPublishSubmitAPIResponse)
}

// ReleaseAlibabaItemPublishSubmitAPIResponse 将 AlibabaItemPublishSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItemPublishSubmitAPIResponse(v *AlibabaItemPublishSubmitAPIResponse) {
	v.Reset()
	poolAlibabaItemPublishSubmitAPIResponse.Put(v)
}
