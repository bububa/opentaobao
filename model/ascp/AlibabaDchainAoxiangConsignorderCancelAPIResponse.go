package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderCancelAPIResponse 自动流转单据取消仓内发货 API返回值
// alibaba.dchain.aoxiang.consignorder.cancel
//
// 自动流转单据取消仓内发货
type AlibabaDchainAoxiangConsignorderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangConsignorderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangConsignorderCancelAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangConsignorderCancelAPIResponseModel is 自动流转单据取消仓内发货 成功返回结果
type AlibabaDchainAoxiangConsignorderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CancelConsignorderResponse *CancelConsignOrderResponse `json:"cancel_consignorder_response,omitempty" xml:"cancel_consignorder_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangConsignorderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CancelConsignorderResponse = nil
}

var poolAlibabaDchainAoxiangConsignorderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangConsignorderCancelAPIResponse)
	},
}

// GetAlibabaDchainAoxiangConsignorderCancelAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderCancelAPIResponse
func GetAlibabaDchainAoxiangConsignorderCancelAPIResponse() *AlibabaDchainAoxiangConsignorderCancelAPIResponse {
	return poolAlibabaDchainAoxiangConsignorderCancelAPIResponse.Get().(*AlibabaDchainAoxiangConsignorderCancelAPIResponse)
}

// ReleaseAlibabaDchainAoxiangConsignorderCancelAPIResponse 将 AlibabaDchainAoxiangConsignorderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderCancelAPIResponse(v *AlibabaDchainAoxiangConsignorderCancelAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderCancelAPIResponse.Put(v)
}
