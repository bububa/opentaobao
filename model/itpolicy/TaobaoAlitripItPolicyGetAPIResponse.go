package itpolicy

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票销售规则】单条查询 API返回值 
taobao.alitrip.it.policy.get

通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
*/
type TaobaoAlitripItPolicyGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripItPolicyGetAPIResponseModel
}

// 【国际机票销售规则】单条查询 成功返回结果
type TaobaoAlitripItPolicyGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_it_policy_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 淘宝政策id
    TaobaoId   int64 `json:"taobao_id,omitempty" xml:"taobao_id,omitempty"`
    // 政策状态，0-未发布，1-已发布，2-已过期
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 扩展字段
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
    // 政策信息
    PolicyDo   *TopPolicyDo `json:"policy_do,omitempty" xml:"policy_do,omitempty"`
}
