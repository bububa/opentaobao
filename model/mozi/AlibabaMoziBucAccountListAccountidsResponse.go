package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一批账号ID查询账号列表 APIResponse
alibaba.mozi.buc.account.list.accountids

根据一批账号ID查询账号列表
*/
type AlibabaMoziBucAccountListAccountidsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziBucAccountListAccountidsResponse
}

type AlibabaMoziBucAccountListAccountidsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_buc_account_list_accountids_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *ListAccountsByAccountIdsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
