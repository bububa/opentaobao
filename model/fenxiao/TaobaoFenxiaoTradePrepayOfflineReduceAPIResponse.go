package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse 渠道分销供应商上传线下流水预存款（减少） API返回值
// taobao.fenxiao.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
type TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoTradePrepayOfflineReduceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoTradePrepayOfflineReduceAPIResponseModel).Reset()
}

// TaobaoFenxiaoTradePrepayOfflineReduceAPIResponseModel is 渠道分销供应商上传线下流水预存款（减少） 成功返回结果
type TaobaoFenxiaoTradePrepayOfflineReduceAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_trade_prepay_offline_reduce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoTradePrepayOfflineReduceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse)
	},
}

// GetTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse
func GetTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse() *TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse {
	return poolTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse.Get().(*TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse)
}

// ReleaseTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse 将 TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse(v *TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoTradePrepayOfflineReduceAPIResponse.Put(v)
}
