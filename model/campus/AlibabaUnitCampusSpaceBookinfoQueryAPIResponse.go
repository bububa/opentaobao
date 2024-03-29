package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaUnitCampusSpaceBookinfoQueryAPIResponse 环路资源信息查询单元环境 API返回值
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
type AlibabaUnitCampusSpaceBookinfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaUnitCampusSpaceBookinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaUnitCampusSpaceBookinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaUnitCampusSpaceBookinfoQueryAPIResponseModel).Reset()
}

// AlibabaUnitCampusSpaceBookinfoQueryAPIResponseModel is 环路资源信息查询单元环境 成功返回结果
type AlibabaUnitCampusSpaceBookinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_unit_campus_space_bookinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 本次调用返回errorCode，一般为错误代码
	ErrorCodeStatus string `json:"error_code_status,omitempty" xml:"error_code_status,omitempty"`
	// 扩展字段
	ErrorExtInfo string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
	// 本次调用返回的消息，一般为错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误level
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 具体返回内容
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 标识本次调用是否执行成功
	SuccessStatus bool `json:"success_status,omitempty" xml:"success_status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaUnitCampusSpaceBookinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCodeStatus = ""
	m.ErrorExtInfo = ""
	m.ErrorMsg = ""
	m.ErrorLevel = ""
	m.Content = nil
	m.SuccessStatus = false
}

var poolAlibabaUnitCampusSpaceBookinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaUnitCampusSpaceBookinfoQueryAPIResponse)
	},
}

// GetAlibabaUnitCampusSpaceBookinfoQueryAPIResponse 从 sync.Pool 获取 AlibabaUnitCampusSpaceBookinfoQueryAPIResponse
func GetAlibabaUnitCampusSpaceBookinfoQueryAPIResponse() *AlibabaUnitCampusSpaceBookinfoQueryAPIResponse {
	return poolAlibabaUnitCampusSpaceBookinfoQueryAPIResponse.Get().(*AlibabaUnitCampusSpaceBookinfoQueryAPIResponse)
}

// ReleaseAlibabaUnitCampusSpaceBookinfoQueryAPIResponse 将 AlibabaUnitCampusSpaceBookinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaUnitCampusSpaceBookinfoQueryAPIResponse(v *AlibabaUnitCampusSpaceBookinfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaUnitCampusSpaceBookinfoQueryAPIResponse.Put(v)
}
