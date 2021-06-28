package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品质量分查询 APIResponse
alibaba.icbu.product.score.get

产品质量分查询
*/
type AlibabaIcbuProductScoreGetAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuProductScoreGetResponse
}

type AlibabaIcbuProductScoreGetResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_product_score_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *ProductScoreInfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
