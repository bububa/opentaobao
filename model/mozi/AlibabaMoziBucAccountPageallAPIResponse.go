package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziBucAccountPageallAPIResponse 查询租户内内所有账号 API返回值
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
type AlibabaMoziBucAccountPageallAPIResponse struct {
	model.CommonResponse
	AlibabaMoziBucAccountPageallAPIResponseModel
}

// AlibabaMoziBucAccountPageallAPIResponseModel is 查询租户内内所有账号 成功返回结果
type AlibabaMoziBucAccountPageallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_buc_account_pageall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageAllAccountsResult `json:"result,omitempty" xml:"result,omitempty"`
}
