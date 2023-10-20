package eleenterpriseordernew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGetstatusAPIResponse 订单状态查询接口 API返回值
// alibaba.ele.enterprise.ordernew.getstatus
//
// 订单状态查询接口
type AlibabaEleEnterpriseOrdernewGetstatusAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGetstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel is 订单状态查询接口 成功返回结果
type AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_getstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回值信息
	EnterpriseData *StandardOrderTrackingInfoDto `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseOrdernewGetstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseOrdernewGetstatusAPIResponse)
	},
}

// GetAlibabaEleEnterpriseOrdernewGetstatusAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGetstatusAPIResponse
func GetAlibabaEleEnterpriseOrdernewGetstatusAPIResponse() *AlibabaEleEnterpriseOrdernewGetstatusAPIResponse {
	return poolAlibabaEleEnterpriseOrdernewGetstatusAPIResponse.Get().(*AlibabaEleEnterpriseOrdernewGetstatusAPIResponse)
}

// ReleaseAlibabaEleEnterpriseOrdernewGetstatusAPIResponse 将 AlibabaEleEnterpriseOrdernewGetstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGetstatusAPIResponse(v *AlibabaEleEnterpriseOrdernewGetstatusAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGetstatusAPIResponse.Put(v)
}
