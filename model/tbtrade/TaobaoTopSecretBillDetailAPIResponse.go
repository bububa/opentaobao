package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretBillDetailAPIResponse 服务商的商家解密账单详情查询 API返回值
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
type TaobaoTopSecretBillDetailAPIResponse struct {
	model.CommonResponse
	TaobaoTopSecretBillDetailAPIResponseModel
}

// TaobaoTopSecretBillDetailAPIResponseModel is 服务商的商家解密账单详情查询 成功返回结果
type TaobaoTopSecretBillDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"top_secret_bill_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账单明细
	Data []BillDetailInfo `json:"data,omitempty" xml:"data>bill_detail_info,omitempty"`
	// 账单总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
