package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销供应商获取清关材料 APIResponse
tmall.hk.clearance.distribution.get

供销体系下，提供供应商可以直接获取其订单身份证信息的接口，以使其完成清关。
*/
type TmallHkClearanceDistributionGetAPIResponse struct {
    model.CommonResponse
    TmallHkClearanceDistributionGetResponse
}

type TmallHkClearanceDistributionGetResponse struct {
    XMLName xml.Name `xml:"tmall_hk_clearance_distribution_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果对象
    
    Result   *CertifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
