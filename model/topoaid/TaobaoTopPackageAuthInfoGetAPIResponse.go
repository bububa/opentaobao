package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageAuthInfoGetAPIResponse 淘宝末端包裹信息获取 API返回值
// taobao.top.package.auth.info.get
//
// 淘宝末端包裹信息获取
type TaobaoTopPackageAuthInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageAuthInfoGetAPIResponseModel
}

// TaobaoTopPackageAuthInfoGetAPIResponseModel is 淘宝末端包裹信息获取 成功返回结果
type TaobaoTopPackageAuthInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_auth_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *IsPrivacyPackageResponse `json:"result,omitempty" xml:"result,omitempty"`
}
