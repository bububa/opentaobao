package mozi

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziBucAccountListAccountidsAPIResponse 根据一批账号ID查询账号列表 API返回值
// alibaba.mozi.buc.account.list.accountids
//
// 根据一批账号ID查询账号列表
type AlibabaMoziBucAccountListAccountidsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziBucAccountListAccountidsAPIResponseModel
}

// AlibabaMoziBucAccountListAccountidsAPIResponseModel is 根据一批账号ID查询账号列表 成功返回结果
type AlibabaMoziBucAccountListAccountidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_buc_account_list_accountids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *ListAccountsByAccountIdsResult `json:"result,omitempty" xml:"result,omitempty"`
}
