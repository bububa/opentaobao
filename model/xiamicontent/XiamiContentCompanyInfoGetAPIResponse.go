package xiamicontent

import (
	"encoding/xml"

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
