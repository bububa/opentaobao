package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse 阿里巴巴.天猫家装.揽配.物流.获取运单号 API返回值
// alibaba.tianmao.lanpei.logistics.mailno
//
// 阿里巴巴.天猫家装.揽配.物流.获取运单号
type AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse struct {
	model.CommonResponse
	AlibabaTianmaoLanpeiLogisticsMailnoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTianmaoLanpeiLogisticsMailnoAPIResponseModel).Reset()
}

// AlibabaTianmaoLanpeiLogisticsMailnoAPIResponseModel is 阿里巴巴.天猫家装.揽配.物流.获取运单号 成功返回结果
type AlibabaTianmaoLanpeiLogisticsMailnoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tianmao_lanpei_logistics_mailno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTianmaoLanpeiLogisticsMailnoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse)
	},
}

// GetAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse 从 sync.Pool 获取 AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse
func GetAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse() *AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse {
	return poolAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse.Get().(*AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse)
}

// ReleaseAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse 将 AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse(v *AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse) {
	v.Reset()
	poolAlibabaTianmaoLanpeiLogisticsMailnoAPIResponse.Put(v)
}
