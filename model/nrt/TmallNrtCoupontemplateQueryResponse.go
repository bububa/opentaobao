package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
券模板查询 APIResponse
tmall.nrt.coupontemplate.query

新零售场景，商家拉取在新零售工作台设置的券数据
*/
type TmallNrtCoupontemplateQueryAPIResponse struct {
    model.CommonResponse
    TmallNrtCoupontemplateQueryResponse
}

type TmallNrtCoupontemplateQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_coupontemplate_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // model
    
    Model   *PageData `json:"model,omitempty" xml:"model,omitempty"`

    
}
