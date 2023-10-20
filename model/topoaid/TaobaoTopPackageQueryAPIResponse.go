package topoaid

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageQueryAPIResponse 淘系包裹查询 API返回值
// taobao.top.package.query
//
// 淘系包裹查询
type TaobaoTopPackageQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageQueryAPIResponseModel
}

// TaobaoTopPackageQueryAPIResponseModel is 淘系包裹查询 成功返回结果
type TaobaoTopPackageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包裹信息列表
	Data []PackageInfo `json:"data,omitempty" xml:"data>package_info,omitempty"`
	// 面单总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
