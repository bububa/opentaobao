package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcgroupsgetAPIResponse 获取自定义用户分组列表 API返回值
// taobao.tmc.groups.get
//
// 获取自定义用户分组列表
type TaobaotmcgroupsgetAPIResponse struct {
	model.CommonResponse
	TaobaotmcgroupsgetAPIResponseModel
}

// TaobaotmcgroupsgetAPIResponseModel is 获取自定义用户分组列表 成功返回结果
type TaobaotmcgroupsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_groups_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dasdasd
	Groups []TmcGroup `json:"groups,omitempty" xml:"groups>tmc_group,omitempty"`
	// 分组总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
