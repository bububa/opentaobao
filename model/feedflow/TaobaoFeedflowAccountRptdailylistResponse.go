package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告主分日数据 APIResponse
taobao.feedflow.account.rptdailylist

获取广告主分日数据
*/
type TaobaoFeedflowAccountRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowAccountRptdailylistResponse
}

type TaobaoFeedflowAccountRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_account_rptdailylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoFeedflowAccountRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
