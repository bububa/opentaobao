package nrpos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse 去前置机商品在线查询 API返回值
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse struct {
	model.CommonResponse
	AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel).Reset()
}

// AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel is 去前置机商品在线查询 成功返回结果
type AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_commdy_posmerchandise_getmerchandise_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse)
	},
}

// GetAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse 从 sync.Pool 获取 AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse
func GetAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse() *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse {
	return poolAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse.Get().(*AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse)
}

// ReleaseAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse 将 AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse(v *AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse) {
	v.Reset()
	poolAlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse.Put(v)
}
