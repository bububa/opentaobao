package nrpos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommdyOfflineGetfileurlAPIResponse 去前置机pos商品离线文件下载地址查询接口 API返回值
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
type AlibabaMosCommdyOfflineGetfileurlAPIResponse struct {
	model.CommonResponse
	AlibabaMosCommdyOfflineGetfileurlAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosCommdyOfflineGetfileurlAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosCommdyOfflineGetfileurlAPIResponseModel).Reset()
}

// AlibabaMosCommdyOfflineGetfileurlAPIResponseModel is 去前置机pos商品离线文件下载地址查询接口 成功返回结果
type AlibabaMosCommdyOfflineGetfileurlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_commdy_offline_getfileurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosCommdyOfflineGetfileurlResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosCommdyOfflineGetfileurlAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosCommdyOfflineGetfileurlAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosCommdyOfflineGetfileurlAPIResponse)
	},
}

// GetAlibabaMosCommdyOfflineGetfileurlAPIResponse 从 sync.Pool 获取 AlibabaMosCommdyOfflineGetfileurlAPIResponse
func GetAlibabaMosCommdyOfflineGetfileurlAPIResponse() *AlibabaMosCommdyOfflineGetfileurlAPIResponse {
	return poolAlibabaMosCommdyOfflineGetfileurlAPIResponse.Get().(*AlibabaMosCommdyOfflineGetfileurlAPIResponse)
}

// ReleaseAlibabaMosCommdyOfflineGetfileurlAPIResponse 将 AlibabaMosCommdyOfflineGetfileurlAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosCommdyOfflineGetfileurlAPIResponse(v *AlibabaMosCommdyOfflineGetfileurlAPIResponse) {
	v.Reset()
	poolAlibabaMosCommdyOfflineGetfileurlAPIResponse.Put(v)
}
