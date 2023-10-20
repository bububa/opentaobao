package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageUnauthQueryAPIResponse 查询某手机号下的包裹 API返回值
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
type TaobaoTopPackageUnauthQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageUnauthQueryAPIResponseModel
}

// TaobaoTopPackageUnauthQueryAPIResponseModel is 查询某手机号下的包裹 成功返回结果
type TaobaoTopPackageUnauthQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_unauth_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包裹信息
	Data []PackageInfo `json:"data,omitempty" xml:"data>package_info,omitempty"`
	// 总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
