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
	RequestId     string         `json:"request_id,omitempty" xml:"qianniu_kefueval_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 评价结果数
    
    ResultCount   int64 `json:"result_count,omitempty" xml:"