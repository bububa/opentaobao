package lsttrade

import (
	"encoding/xml"

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

// AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel is 卖家退款单商品状态同步 成功返回结果
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_fastrefund_goodsstatus_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示成功，false表示失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
