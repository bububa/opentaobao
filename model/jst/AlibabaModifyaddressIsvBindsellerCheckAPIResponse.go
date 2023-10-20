package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaModifyaddressIsvBindsellerCheckAPIResponse 查询服务商下的商家是否开通了改地址 API返回值
// alibaba.modifyaddress.isv.bindseller.check
//
// 鉴权服务商和商家的绑定关系，看商家是否开通了改地址
// 1. 没有授权
// 2. 已与当前appkey签约
// 3. 没有签约
// 4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
type AlibabaModifyaddressIsvBindsellerCheckAPIResponse struct {
	model.CommonResponse
	AlibabaModifyaddressIsvBindsellerCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaModifyaddressIsvBindsellerCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaModifyaddressIsvBindsellerCheckAPIResponseModel).Reset()
}

// AlibabaModifyaddressIsvBindsellerCheckAPIResponseModel is 查询服务商下的商家是否开通了改地址 成功返回结果
type AlibabaModifyaddressIsvBindsellerCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_modifyaddress_isv_bindseller_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model *CheckSellerChooseErpResponse `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaModifyaddressIsvBindsellerCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolAlibabaModifyaddressIsvBindsellerCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaModifyaddressIsvBindsellerCheckAPIResponse)
	},
}

// GetAlibabaModifyaddressIsvBindsellerCheckAPIResponse 从 sync.Pool 获取 AlibabaModifyaddressIsvBindsellerCheckAPIResponse
func GetAlibabaModifyaddressIsvBindsellerCheckAPIResponse() *AlibabaModifyaddressIsvBindsellerCheckAPIResponse {
	return poolAlibabaModifyaddressIsvBindsellerCheckAPIResponse.Get().(*AlibabaModifyaddressIsvBindsellerCheckAPIResponse)
}

// ReleaseAlibabaModifyaddressIsvBindsellerCheckAPIResponse 将 AlibabaModifyaddressIsvBindsellerCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaModifyaddressIsvBindsellerCheckAPIResponse(v *AlibabaModifyaddressIsvBindsellerCheckAPIResponse) {
	v.Reset()
	poolAlibabaModifyaddressIsvBindsellerCheckAPIResponse.Put(v)
}
