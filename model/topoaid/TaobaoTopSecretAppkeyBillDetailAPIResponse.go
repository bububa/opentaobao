package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretAppkeyBillDetailAPIResponse 服务商解密账单查询 API返回值
// taobao.top.secret.appkey.bill.detail
//
// 服务商解密账单查询,分页返回所有店铺的账单，每个店铺每天仅包含两条数据，当天产生的号租费 和 当天产生的通话费，仅对90天内的账单提供SLA保障。查询账单详情请使用taobao.top.secret.bill.detail接口。
type TaobaoTopSecretAppkeyBillDetailAPIResponse struct {
	model.CommonResponse
	TaobaoTopSecretAppkeyBillDetailAPIResponseModel
}

// TaobaoTopSecretAppkeyBillDetailAPIResponseModel is 服务商解密账单查询 成功返回结果
type TaobaoTopSecretAppkeyBillDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"top_secret_appkey_bill_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 账单明细
	Data []BillDetailInfo `json:"data,omitempty" xml:"data>bill_detail_info,omitempty"`
	// 账单总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
