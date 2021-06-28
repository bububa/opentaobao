package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
类目预测接口 APIResponse
alibaba.imap.category.predict

* 类目预测接口
     * 【必填字段】 title, srcChannelId, srcCategoryId, targetChannelId
     * 【非必填，但有最好填上】itemId, barcode, brandName, pvPairDOList, srcCatNamePathList
*/
type AlibabaImapCategoryPredictAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaImapCategoryPredictResponse `json:"alibaba_imap_category_predict_response,omitempty"` 
    AlibabaImapCategoryPredictResponse
}

/* model for simplify = false
type AlibabaImapCategoryPredictResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaImapCategoryPredictResult  *AlibabaImapCategoryPredictResult `json:"alibaba_imap_category_predict_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaImapCategoryPredictResponse struct {

    // 接口返回model
    Result   *AlibabaImapCategoryPredictResult `json:"result,omitempty"`

}
