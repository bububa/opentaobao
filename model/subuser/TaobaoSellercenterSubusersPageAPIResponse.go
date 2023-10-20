package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercentersubuserspageAPIResponse 通过主账号登陆态分页查询子账号列表 API返回值
// taobao.sellercenter.subusers.page
//
// 通过主账号登陆态分页查询子账号列表
type TaobaosellercentersubuserspageAPIResponse struct {
	model.CommonResponse
	TaobaosellercentersubuserspageAPIResponseModel
}

// TaobaosellercentersubuserspageAPIResponseModel is 通过主账号登陆态分页查询子账号列表 成功返回结果
type TaobaosellercentersubuserspageAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_subusers_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流
	Subusers []SubUserInfo `json:"subusers,omitempty" xml:"subusers>sub_user_info,omitempty"`
	// isv本次调用传入的页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 本次调用删选条件下的总子账号数量，除以页大小可得出最大页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// isv本次调用传入的页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
