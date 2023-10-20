package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse 预售商家仓接单 API返回值
// alibaba.ascp.uop.taobao.presalesorder.create
//
// 预售商家仓接单
type AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopTaobaoPresalesorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopTaobaoPresalesorderCreateAPIResponseModel).Reset()
}

// AlibabaAscpUopTaobaoPresalesorderCreateAPIResponseModel is 预售商家仓接单 成功返回结果
type AlibabaAscpUopTaobaoPresalesorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_taobao_presalesorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	PresalesOrderCreateResponse *ResultWrapper `json:"presales_order_create_response,omitempty" xml:"presales_order_create_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopTaobaoPresalesorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PresalesOrderCreateResponse = nil
}

var poolAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse)
	},
}

// GetAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse 从 sync.Pool 获取 AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse
func GetAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse() *AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse {
	return poolAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse.Get().(*AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse)
}

// ReleaseAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse 将 AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse(v *AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopTaobaoPresalesorderCreateAPIResponse.Put(v)
}
