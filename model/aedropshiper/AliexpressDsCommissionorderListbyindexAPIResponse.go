package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsCommissionorderListbyindexAPIResponse 联盟订单分页查询 API返回值
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
type AliexpressDsCommissionorderListbyindexAPIResponse struct {
	model.CommonResponse
	AliexpressDsCommissionorderListbyindexAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressDsCommissionorderListbyindexAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressDsCommissionorderListbyindexAPIResponseModel).Reset()
}

// AliexpressDsCommissionorderListbyindexAPIResponseModel is 联盟订单分页查询 成功返回结果
type AliexpressDsCommissionorderListbyindexAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_commissionorder_listbyindex_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// System Error
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// result object
	Result *TrafficOrderResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressDsCommissionorderListbyindexAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspMsg = ""
	m.RspCode = ""
	m.Result = nil
}

var poolAliexpressDsCommissionorderListbyindexAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressDsCommissionorderListbyindexAPIResponse)
	},
}

// GetAliexpressDsCommissionorderListbyindexAPIResponse 从 sync.Pool 获取 AliexpressDsCommissionorderListbyindexAPIResponse
func GetAliexpressDsCommissionorderListbyindexAPIResponse() *AliexpressDsCommissionorderListbyindexAPIResponse {
	return poolAliexpressDsCommissionorderListbyindexAPIResponse.Get().(*AliexpressDsCommissionorderListbyindexAPIResponse)
}

// ReleaseAliexpressDsCommissionorderListbyindexAPIResponse 将 AliexpressDsCommissionorderListbyindexAPIResponse 保存到 sync.Pool
func ReleaseAliexpressDsCommissionorderListbyindexAPIResponse(v *AliexpressDsCommissionorderListbyindexAPIResponse) {
	v.Reset()
	poolAliexpressDsCommissionorderListbyindexAPIResponse.Put(v)
}
