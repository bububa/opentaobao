package ticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketSkusBatchUploadAPIResponse 【门票API2.0】门票价格库存同步接口（多票种批量更新） API返回值
// alitrip.ticket.skus.batch.upload
//
// 飞猪度假新版门票商品价格库存同步接口（多票种批量更新）。
// 注1、一个票种下可以挂多个规则（规则id必须不一样，每个规则实际对应了一个sku），同一个规则可以在不同票种下使用。
// 注2、日历库存和区间库存门票，统一使用DateInventory结构。对于日历库存门票请上传每一天的价格库存；对于区间库存门票，建议只上传开始和结束日期的价格库存，也支持上传每天的价格库存，系统会自动进行聚合（取第一天的价格为区间价格，累计所有天的库存为区间库存）。
// 注3、该接口同时支持 新增某个规则的价格库存 和 更新现有规则的价格库存。如果不清楚是否已在某个规则下上传过价格库存，请使用alitrip.ticket.product.query接口进行查询。如果该规则在该票种下已经存在，则该接口会判断为是价格库存更新操作。
type AlitripTicketSkusBatchUploadAPIResponse struct {
	model.CommonResponse
	AlitripTicketSkusBatchUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTicketSkusBatchUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTicketSkusBatchUploadAPIResponseModel).Reset()
}

// AlitripTicketSkusBatchUploadAPIResponseModel is 【门票API2.0】门票价格库存同步接口（多票种批量更新） 成功返回结果
type AlitripTicketSkusBatchUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ticket_skus_batch_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 价格库存同步结果
	UpdateResult *TicketItemResult `json:"update_result,omitempty" xml:"update_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTicketSkusBatchUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateResult = nil
}

var poolAlitripTicketSkusBatchUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTicketSkusBatchUploadAPIResponse)
	},
}

// GetAlitripTicketSkusBatchUploadAPIResponse 从 sync.Pool 获取 AlitripTicketSkusBatchUploadAPIResponse
func GetAlitripTicketSkusBatchUploadAPIResponse() *AlitripTicketSkusBatchUploadAPIResponse {
	return poolAlitripTicketSkusBatchUploadAPIResponse.Get().(*AlitripTicketSkusBatchUploadAPIResponse)
}

// ReleaseAlitripTicketSkusBatchUploadAPIResponse 将 AlitripTicketSkusBatchUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripTicketSkusBatchUploadAPIResponse(v *AlitripTicketSkusBatchUploadAPIResponse) {
	v.Reset()
	poolAlitripTicketSkusBatchUploadAPIResponse.Put(v)
}
