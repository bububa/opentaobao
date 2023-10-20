package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse 预售商家仓出库 API返回值
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponseModel).Reset()
}

// AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponseModel is 预售商家仓出库 成功返回结果
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_taobao_presalesorder_consignconfirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	PresalesOrderConsignConfirmResponse *ResultWrapper `json:"presales_order_consign_confirm_response,omitempty" xml:"presales_order_consign_confirm_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PresalesOrderConsignConfirmResponse = nil
}

var poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse)
	},
}

// GetAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse 从 sync.Pool 获取 AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse
func GetAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse() *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse {
	return poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse.Get().(*AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse)
}

// ReleaseAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse 将 AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse(v *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse.Put(v)
}
