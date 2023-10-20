package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse 卖家退款单商品状态同步 API返回值
// alibaba.lst.trade.fastrefund.goodsstatus.sync
//
// 卖家退款单商品状态同步
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel).Reset()
}

// AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel is 卖家退款单商品状态同步 成功返回结果
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_fastrefund_goodsstatus_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示成功，false表示失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse)
	},
}

// GetAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse 从 sync.Pool 获取 AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse
func GetAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse() *AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse {
	return poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse.Get().(*AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse)
}

// ReleaseAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse 将 AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse(v *AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse.Put(v)
}
