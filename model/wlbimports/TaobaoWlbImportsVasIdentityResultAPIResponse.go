package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsvasidentityresultAPIResponse 集货鉴定结果 API返回值
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
type TaobaowlbimportsvasidentityresultAPIResponse struct {
	model.CommonResponse
	TaobaowlbimportsvasidentityresultAPIResponseModel
}

// TaobaowlbimportsvasidentityresultAPIResponseModel is 集货鉴定结果 成功返回结果
type TaobaowlbimportsvasidentityresultAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_vas_identity_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回出参数结果
	Result *TaobaowlbimportsvasidentityresultTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
