package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoTradePrepayOfflineAddAPIResponse 线下预存款流水增加 API返回值
// taobao.fenxiao.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaoFenxiaoTradePrepayOfflineAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoTradePrepayOfflineAddAPIResponseModel
}

// TaobaoFenxiaoTradePrepayOfflineAddAPIResponseModel is 线下预存款流水增加 成功返回结果
type TaobaoFenxiaoTradePrepayOfflineAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_trade_prepay_offline_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoFenxiaoTradePrepayOfflineAddResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}
