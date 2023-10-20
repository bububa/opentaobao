package ott

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttAlicbFacadeserviceGetdataAPIResponse 影视SDK获取设备能力值 API返回值
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
type YoukuOttAlicbFacadeserviceGetdataAPIResponse struct {
	model.CommonResponse
	YoukuOttAlicbFacadeserviceGetdataAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttAlicbFacadeserviceGetdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttAlicbFacadeserviceGetdataAPIResponseModel).Reset()
}

// YoukuOttAlicbFacadeserviceGetdataAPIResponseModel is 影视SDK获取设备能力值 成功返回结果
type YoukuOttAlicbFacadeserviceGetdataAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_alicb_facadeservice_getdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备能力JSON
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttAlicbFacadeserviceGetdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
}

var poolYoukuOttAlicbFacadeserviceGetdataAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttAlicbFacadeserviceGetdataAPIResponse)
	},
}

// GetYoukuOttAlicbFacadeserviceGetdataAPIResponse 从 sync.Pool 获取 YoukuOttAlicbFacadeserviceGetdataAPIResponse
func GetYoukuOttAlicbFacadeserviceGetdataAPIResponse() *YoukuOttAlicbFacadeserviceGetdataAPIResponse {
	return poolYoukuOttAlicbFacadeserviceGetdataAPIResponse.Get().(*YoukuOttAlicbFacadeserviceGetdataAPIResponse)
}

// ReleaseYoukuOttAlicbFacadeserviceGetdataAPIResponse 将 YoukuOttAlicbFacadeserviceGetdataAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttAlicbFacadeserviceGetdataAPIResponse(v *YoukuOttAlicbFacadeserviceGetdataAPIResponse) {
	v.Reset()
	poolYoukuOttAlicbFacadeserviceGetdataAPIResponse.Put(v)
}
