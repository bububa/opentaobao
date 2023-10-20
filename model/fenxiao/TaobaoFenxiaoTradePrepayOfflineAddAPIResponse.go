package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaotradeprepayofflineaddAPIResponse 线下预存款流水增加 API返回值
// taobao.fenxiao.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaofenxiaotradeprepayofflineaddAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaotradeprepayofflineaddAPIResponseModel
}

// TaobaofenxiaotradeprepayofflineaddAPIResponseModel is 线下预存款流水增加 成功返回结果
type TaobaofenxiaotradeprepayofflineaddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_trade_prepay_offline_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaofenxiaotradeprepayofflineaddResultTopDo `json:"result,omitempty" xml:"result,omitempty"`
}
