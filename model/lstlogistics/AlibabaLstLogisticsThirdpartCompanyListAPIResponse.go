package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsThirdpartCompanyListAPIResponse 供应商-异云-第三方物流公司列表 API返回值
// alibaba.lst.logistics.thirdpart.company.list
//
// 异地云仓发货时，需填写的第三方物流公司列表
type AlibabaLstLogisticsThirdpartCompanyListAPIResponse struct {
	model.CommonResponse
	AlibabaLstLogisticsThirdpartCompanyListAPIResponseModel
}

// AlibabaLstLogisticsThirdpartCompanyListAPIResponseModel is 供应商-异云-第三方物流公司列表 成功返回结果
type AlibabaLstLogisticsThirdpartCompanyListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_thirdpart_company_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回内容
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
}
