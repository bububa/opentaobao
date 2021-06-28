package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客服评价详情接口 APIResponse
taobao.qianniu.kefueval.get

获取买家对客服的服务评价
*/
type TaobaoQianniuKefuevalGetAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuKefuevalGetResponse
}

type TaobaoQianniuKefuevalGetResponse struct {
    XMLName xml.Name `xml:"qianniu_kefueval_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 评价结果数
    
    ResultCount   int64 `json:"result_count,omitempty" xml:"result_count,omitempty"`

    
    // 评价明细
    
    StaffEvalDetails   []EvalDetail `json:"staff_eval_details,omitempty" xml:"staff_eval_details>eval_detail,omitempty"`
    
    
}
