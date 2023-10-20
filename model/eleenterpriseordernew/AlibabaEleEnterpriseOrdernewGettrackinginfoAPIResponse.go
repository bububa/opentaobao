package eleenterpriseordernew

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse 订单配送信息跟踪 API返回值
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel).Reset()
}

// AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel is 订单配送信息跟踪 成功返回结果
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_gettrackinginfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
	// 返回信息
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EnterpriseCode = ""
	m.EnterpriseMsg = ""
	m.EnterpriseRequestid = ""
	m.EnterpriseData = nil
}

var poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse)
	},
}

// GetAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse
func GetAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse() *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse {
	return poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse.Get().(*AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse)
}

// ReleaseAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse 将 AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse(v *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse.Put(v)
}
