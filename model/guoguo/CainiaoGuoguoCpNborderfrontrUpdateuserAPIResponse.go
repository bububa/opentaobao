package guoguo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse 小件员信息变更 API返回值
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse struct {
	model.CommonResponse
	CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponseModel).Reset()
}

// CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponseModel is 小件员信息变更 成功返回结果
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_cp_nborderfrontr_updateuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// errorCode
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StatusMessage = ""
	m.StatusCode = ""
	m.IsSuccess = false
}

var poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse)
	},
}

// GetCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse 从 sync.Pool 获取 CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse
func GetCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse() *CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse {
	return poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse.Get().(*CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse)
}

// ReleaseCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse 将 CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse(v *CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse) {
	v.Reset()
	poolCainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse.Put(v)
}
