package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundGoodsWaybillAPIResponse 淘外分销退货回传物流单号 API返回值
// alibaba.ascp.channel.refund.goods.waybill
//
// 淘外分销退货回传物流单号
type AlibabaAscpChannelRefundGoodsWaybillAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelRefundGoodsWaybillAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundGoodsWaybillAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelRefundGoodsWaybillAPIResponseModel).Reset()
}

// AlibabaAscpChannelRefundGoodsWaybillAPIResponseModel is 淘外分销退货回传物流单号 成功返回结果
type AlibabaAscpChannelRefundGoodsWaybillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_refund_goods_waybill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpChannelRefundGoodsWaybillResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundGoodsWaybillAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelRefundGoodsWaybillAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelRefundGoodsWaybillAPIResponse)
	},
}

// GetAlibabaAscpChannelRefundGoodsWaybillAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelRefundGoodsWaybillAPIResponse
func GetAlibabaAscpChannelRefundGoodsWaybillAPIResponse() *AlibabaAscpChannelRefundGoodsWaybillAPIResponse {
	return poolAlibabaAscpChannelRefundGoodsWaybillAPIResponse.Get().(*AlibabaAscpChannelRefundGoodsWaybillAPIResponse)
}

// ReleaseAlibabaAscpChannelRefundGoodsWaybillAPIResponse 将 AlibabaAscpChannelRefundGoodsWaybillAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelRefundGoodsWaybillAPIResponse(v *AlibabaAscpChannelRefundGoodsWaybillAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelRefundGoodsWaybillAPIResponse.Put(v)
}
