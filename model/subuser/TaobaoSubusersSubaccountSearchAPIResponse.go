package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersSubaccountSearchAPIResponse 根据子账号登录名后缀模糊搜索子账号列表 API返回值
// taobao.subusers.subaccount.search
//
// 根据子账号冒号后缀搜索子账号列表，支持中文单字、英文单词（不支持英文单字母） 分词规则搜索，该搜索词必传。模糊搜索使用阿里云搜索引擎所以该接口增值收费，如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口
type TaobaoSubusersSubaccountSearchAPIResponse struct {
	model.CommonResponse
	TaobaoSubusersSubaccountSearchAPIResponseModel
}

// TaobaoSubusersSubaccountSearchAPIResponseModel is 根据子账号登录名后缀模糊搜索子账号列表 成功返回结果
type TaobaoSubusersSubaccountSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"subusers_subaccount_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号列表
	Subaccounts []SubAccountInfo `json:"subaccounts,omitempty" xml:"subaccounts>sub_account_info,omitempty"`
	// isv本次调用传入的页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 本次调用删选条件下的总子账号数量，除以页大小可得出最大页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// isv本次调用传入的页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
