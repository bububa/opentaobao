package cainiaoncwl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoncwljhorderqueryAPIResponse 农村物流集货单查询接口 API返回值
// cainiao.ncwl.jhorder.query
//
// 提供给接入商家，查询农村物流集货单
type CainiaoncwljhorderqueryAPIResponse struct {
	model.CommonResponse
	CainiaoncwljhorderqueryAPIResponseModel
}

// CainiaoncwljhorderqueryAPIResponseModel is 农村物流集货单查询接口 成功返回结果
type CainiaoncwljhorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_ncwl_jhorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询获得的集货单列表
	JhOrderList []JhOrder `json:"jh_order_list,omitempty" xml:"jh_order_list>jh_order,omitempty"`
	// 总结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 翻页查询当前返回页数；指定集货号时无效；
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
