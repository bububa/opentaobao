package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportParterSynccardAPIResponse 阿里体育-卡信息同步接口 API返回值
// alibaba.alisports.passport.parter.synccard
//
// 运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
type AlibabaAlisportsPassportParterSynccardAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportParterSynccardAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportParterSynccardAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportParterSynccardAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportParterSynccardAPIResponseModel is 阿里体育-卡信息同步接口 成功返回结果
type AlibabaAlisportsPassportParterSynccardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_parter_synccard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 正确或错误的信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 返回数据
	AlispData string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 200标识成功，其他的code为失败
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportParterSynccardAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispData = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsPassportParterSynccardAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportParterSynccardAPIResponse)
	},
}

// GetAlibabaAlisportsPassportParterSynccardAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportParterSynccardAPIResponse
func GetAlibabaAlisportsPassportParterSynccardAPIResponse() *AlibabaAlisportsPassportParterSynccardAPIResponse {
	return poolAlibabaAlisportsPassportParterSynccardAPIResponse.Get().(*AlibabaAlisportsPassportParterSynccardAPIResponse)
}

// ReleaseAlibabaAlisportsPassportParterSynccardAPIResponse 将 AlibabaAlisportsPassportParterSynccardAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportParterSynccardAPIResponse(v *AlibabaAlisportsPassportParterSynccardAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportParterSynccardAPIResponse.Put(v)
}
