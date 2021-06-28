package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品质量分查询 APIResponse
alibaba.icbu.product.score.get

产品质量分查询
*/
type AlibabaIcbuProductScoreGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaIcbuProductScoreGetResponse `json:"alibaba_icbu_product_score_get_response,omitempty"` 
    AlibabaIcbuProductScoreGetResponse
}

/* model for simplify = false
type AlibabaIcbuProductScoreGetResponse struct {

    // 系统自动生成
    
    Result  *struct {
        ProductScoreInfoResult  *ProductScoreInfoResult `json:"product_score_info_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaIcbuProductScoreGetResponse struct {

    // 系统自动生成
    Result   *ProductScoreInfoResult `json:"result,omitempty"`

}
