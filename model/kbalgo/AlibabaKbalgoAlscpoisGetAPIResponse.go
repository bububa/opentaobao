package kbalgo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKbalgoAlscpoisGetAPIResponse 百度批量获取本地poi接口 API返回值
// alibaba.kbalgo.alscpois.get
//
// 接口用于百度方获取本地生活poi数据，分页获取。
type AlibabaKbalgoAlscpoisGetAPIResponse struct {
	model.CommonResponse
	AlibabaKbalgoAlscpoisGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKbalgoAlscpoisGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKbalgoAlscpoisGetAPIResponseModel).Reset()
}

// AlibabaKbalgoAlscpoisGetAPIResponseModel is 百度批量获取本地poi接口 成功返回结果
type AlibabaKbalgoAlscpoisGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kbalgo_alscpois_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果根节点。
	Result *AlscPoiToBaiduResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKbalgoAlscpoisGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKbalgoAlscpoisGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKbalgoAlscpoisGetAPIResponse)
	},
}

// GetAlibabaKbalgoAlscpoisGetAPIResponse 从 sync.Pool 获取 AlibabaKbalgoAlscpoisGetAPIResponse
func GetAlibabaKbalgoAlscpoisGetAPIResponse() *AlibabaKbalgoAlscpoisGetAPIResponse {
	return poolAlibabaKbalgoAlscpoisGetAPIResponse.Get().(*AlibabaKbalgoAlscpoisGetAPIResponse)
}

// ReleaseAlibabaKbalgoAlscpoisGetAPIResponse 将 AlibabaKbalgoAlscpoisGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKbalgoAlscpoisGetAPIResponse(v *AlibabaKbalgoAlscpoisGetAPIResponse) {
	v.Reset()
	poolAlibabaKbalgoAlscpoisGetAPIResponse.Put(v)
}
