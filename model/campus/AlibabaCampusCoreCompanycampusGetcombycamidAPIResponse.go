package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse 根据园区ID获取运营公司信息 API返回值
// alibaba.campus.core.companycampus.getcombycamid
//
// 根据园区ID获取运营公司信息
type AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel).Reset()
}

// AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel is 根据园区ID获取运营公司信息 成功返回结果
type AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_core_companycampus_getcombycamid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusCoreCompanycampusGetcombycamidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse)
	},
}

// GetAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse 从 sync.Pool 获取 AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse
func GetAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse() *AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse {
	return poolAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse.Get().(*AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse)
}

// ReleaseAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse 将 AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse(v *AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse) {
	v.Reset()
	poolAlibabaCampusCoreCompanycampusGetcombycamidAPIResponse.Put(v)
}
