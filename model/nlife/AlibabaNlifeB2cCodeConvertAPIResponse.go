package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cCodeConvertAPIResponse b2c转码 API返回值
// alibaba.nlife.b2c.code.convert
//
// 将商品的URL转码，ISV将该码写入RFID
type AlibabaNlifeB2cCodeConvertAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cCodeConvertAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cCodeConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cCodeConvertAPIResponseModel).Reset()
}

// AlibabaNlifeB2cCodeConvertAPIResponseModel is b2c转码 成功返回结果
type AlibabaNlifeB2cCodeConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_code_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *ItemCodeConvertResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cCodeConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaNlifeB2cCodeConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cCodeConvertAPIResponse)
	},
}

// GetAlibabaNlifeB2cCodeConvertAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cCodeConvertAPIResponse
func GetAlibabaNlifeB2cCodeConvertAPIResponse() *AlibabaNlifeB2cCodeConvertAPIResponse {
	return poolAlibabaNlifeB2cCodeConvertAPIResponse.Get().(*AlibabaNlifeB2cCodeConvertAPIResponse)
}

// ReleaseAlibabaNlifeB2cCodeConvertAPIResponse 将 AlibabaNlifeB2cCodeConvertAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cCodeConvertAPIResponse(v *AlibabaNlifeB2cCodeConvertAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cCodeConvertAPIResponse.Put(v)
}
