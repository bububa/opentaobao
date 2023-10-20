package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubuserspageAPIResponse 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口） API返回值
// taobao.subusers.page
//
// 分页获取指定账户的子账号简易信息列表
// （新isv接入建议使用taobao.sellercenter.subusers.page接口）
type TaobaosubuserspageAPIResponse struct {
	model.CommonResponse
	TaobaosubuserspageAPIResponseModel
}

// TaobaosubuserspageAPIResponseModel is 分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口） 成功返回结果
type TaobaosubuserspageAPIResponseModel struct {
	XMLName xml.Name `xml:"subusers_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号基本信息
	Subaccounts []SubAccountInfo `json:"subaccounts,omitempty" xml:"subaccounts>sub_account_info,omitempty"`
	// isv本次调用传入的页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 本次调用删选条件下的总子账号数量，除以页大小可得出最大页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// isv本次调用传入的页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
