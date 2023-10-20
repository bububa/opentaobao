package xiamicontent

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentCompanyInfoGetAPIResponse 获取厂牌信息 API返回值
// xiami.content.company.info.get
//
// 获取厂牌信息
type XiamiContentCompanyInfoGetAPIResponse struct {
	model.CommonResponse
	XiamiContentCompanyInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *XiamiContentCompanyInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.XiamiContentCompanyInfoGetAPIResponseModel).Reset()
}

// XiamiContentCompanyInfoGetAPIResponseModel is 获取厂牌信息 成功返回结果
type XiamiContentCompanyInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xiami_content_company_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 厂牌信息
	CompanyList []CompanyDto `json:"company_list,omitempty" xml:"company_list>company_dto,omitempty"`
	// 结果
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *XiamiContentCompanyInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CompanyList = m.CompanyList[:0]
	m.ResultCode = nil
}

var poolXiamiContentCompanyInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(XiamiContentCompanyInfoGetAPIResponse)
	},
}

// GetXiamiContentCompanyInfoGetAPIResponse 从 sync.Pool 获取 XiamiContentCompanyInfoGetAPIResponse
func GetXiamiContentCompanyInfoGetAPIResponse() *XiamiContentCompanyInfoGetAPIResponse {
	return poolXiamiContentCompanyInfoGetAPIResponse.Get().(*XiamiContentCompanyInfoGetAPIResponse)
}

// ReleaseXiamiContentCompanyInfoGetAPIResponse 将 XiamiContentCompanyInfoGetAPIResponse 保存到 sync.Pool
func ReleaseXiamiContentCompanyInfoGetAPIResponse(v *XiamiContentCompanyInfoGetAPIResponse) {
	v.Reset()
	poolXiamiContentCompanyInfoGetAPIResponse.Put(v)
}
