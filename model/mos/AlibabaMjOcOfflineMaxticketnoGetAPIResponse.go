package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOfflineMaxticketnoGetAPIResponse pos机获取线下最大小票号 API返回值
// alibaba.mj.oc.offline.maxticketno.get
//
// 给pos机提供线下最大小票号查询
type AlibabaMjOcOfflineMaxticketnoGetAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcOfflineMaxticketnoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjOcOfflineMaxticketnoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjOcOfflineMaxticketnoGetAPIResponseModel).Reset()
}

// AlibabaMjOcOfflineMaxticketnoGetAPIResponseModel is pos机获取线下最大小票号 成功返回结果
type AlibabaMjOcOfflineMaxticketnoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_offline_maxticketno_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 时间
	PayDate string `json:"pay_date,omitempty" xml:"pay_date,omitempty"`
	// 收银机号
	PosNo string `json:"pos_no,omitempty" xml:"pos_no,omitempty"`
	// 小票号
	PayNo string `json:"pay_no,omitempty" xml:"pay_no,omitempty"`
	// 外部门店号
	StoreNo string `json:"store_no,omitempty" xml:"store_no,omitempty"`
	// 联合收银标记
	Union bool `json:"union,omitempty" xml:"union,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjOcOfflineMaxticketnoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PayDate = ""
	m.PosNo = ""
	m.PayNo = ""
	m.StoreNo = ""
	m.Union = false
}

var poolAlibabaMjOcOfflineMaxticketnoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcOfflineMaxticketnoGetAPIResponse)
	},
}

// GetAlibabaMjOcOfflineMaxticketnoGetAPIResponse 从 sync.Pool 获取 AlibabaMjOcOfflineMaxticketnoGetAPIResponse
func GetAlibabaMjOcOfflineMaxticketnoGetAPIResponse() *AlibabaMjOcOfflineMaxticketnoGetAPIResponse {
	return poolAlibabaMjOcOfflineMaxticketnoGetAPIResponse.Get().(*AlibabaMjOcOfflineMaxticketnoGetAPIResponse)
}

// ReleaseAlibabaMjOcOfflineMaxticketnoGetAPIResponse 将 AlibabaMjOcOfflineMaxticketnoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjOcOfflineMaxticketnoGetAPIResponse(v *AlibabaMjOcOfflineMaxticketnoGetAPIResponse) {
	v.Reset()
	poolAlibabaMjOcOfflineMaxticketnoGetAPIResponse.Put(v)
}
