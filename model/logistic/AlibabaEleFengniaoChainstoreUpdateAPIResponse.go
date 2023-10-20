package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoChainstoreUpdateAPIResponse 修改门店信息接口 API返回值
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
type AlibabaEleFengniaoChainstoreUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoChainstoreUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoChainstoreUpdateAPIResponseModel).Reset()
}

// AlibabaEleFengniaoChainstoreUpdateAPIResponseModel is 修改门店信息接口 成功返回结果
type AlibabaEleFengniaoChainstoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoChainstoreUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
}

var poolAlibabaEleFengniaoChainstoreUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoChainstoreUpdateAPIResponse)
	},
}

// GetAlibabaEleFengniaoChainstoreUpdateAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoChainstoreUpdateAPIResponse
func GetAlibabaEleFengniaoChainstoreUpdateAPIResponse() *AlibabaEleFengniaoChainstoreUpdateAPIResponse {
	return poolAlibabaEleFengniaoChainstoreUpdateAPIResponse.Get().(*AlibabaEleFengniaoChainstoreUpdateAPIResponse)
}

// ReleaseAlibabaEleFengniaoChainstoreUpdateAPIResponse 将 AlibabaEleFengniaoChainstoreUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoChainstoreUpdateAPIResponse(v *AlibabaEleFengniaoChainstoreUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoChainstoreUpdateAPIResponse.Put(v)
}
