package fivee

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFiveeCompanyGetAPIResponse 查询商信息 API返回值
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
type TaobaoFiveeCompanyGetAPIResponse struct {
	model.CommonResponse
	TaobaoFiveeCompanyGetAPIResponseModel
}

// TaobaoFiveeCompanyGetAPIResponseModel is 查询商信息 成功返回结果
type TaobaoFiveeCompanyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fivee_company_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFiveeCompanyGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
