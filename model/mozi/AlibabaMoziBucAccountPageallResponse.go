package mozi

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租户内内所有账号 APIResponse
alibaba.mozi.buc.account.pageall

查询租户内内所有账号
*/
type AlibabaMoziBucAccountPageallAPIResponse struct {
    model.CommonResponse
    AlibabaMoziBucAccountPageallResponse
}

type AlibabaMoziBucAccountPageallResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_buc_account_pageall_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *PageAllAccountsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
