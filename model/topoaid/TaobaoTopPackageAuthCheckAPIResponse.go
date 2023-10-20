package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotoppackageauthcheckAPIResponse 校验用户授权关系 API返回值
// taobao.top.package.auth.check
//
// 校验用户授权关系
type TaobaotoppackageauthcheckAPIResponse struct {
	model.CommonResponse
	TaobaotoppackageauthcheckAPIResponseModel
}

// TaobaotoppackageauthcheckAPIResponseModel is 校验用户授权关系 成功返回结果
type TaobaotoppackageauthcheckAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_auth_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权查询结果
	Result *AuthScopeCheckResponse `json:"result,omitempty" xml:"result,omitempty"`
}
