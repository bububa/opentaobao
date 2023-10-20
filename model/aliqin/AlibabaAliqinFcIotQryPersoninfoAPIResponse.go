package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotQryPersoninfoAPIResponse 查询物联卡个人实人认证信息 API返回值
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
type AlibabaAliqinFcIotQryPersoninfoAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotQryPersoninfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotQryPersoninfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotQryPersoninfoAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotQryPersoninfoAPIResponseModel is 查询物联卡个人实人认证信息 成功返回结果
type AlibabaAliqinFcIotQryPersoninfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qry_personinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotQryPersoninfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotQryPersoninfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotQryPersoninfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotQryPersoninfoAPIResponse)
	},
}

// GetAlibabaAliqinFcIotQryPersoninfoAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotQryPersoninfoAPIResponse
func GetAlibabaAliqinFcIotQryPersoninfoAPIResponse() *AlibabaAliqinFcIotQryPersoninfoAPIResponse {
	return poolAlibabaAliqinFcIotQryPersoninfoAPIResponse.Get().(*AlibabaAliqinFcIotQryPersoninfoAPIResponse)
}

// ReleaseAlibabaAliqinFcIotQryPersoninfoAPIResponse 将 AlibabaAliqinFcIotQryPersoninfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotQryPersoninfoAPIResponse(v *AlibabaAliqinFcIotQryPersoninfoAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotQryPersoninfoAPIResponse.Put(v)
}
