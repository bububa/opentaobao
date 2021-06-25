package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单量预测 APIResponse
alibaba.wdk.scm.lrp.order.predict

提供基于门店和配送站的订单量预测，可用于提前安排人力资源
*/
type AlibabaWdkScmLrpOrderPredictAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkScmLrpOrderPredictResponse `json:"alibaba_wdk_scm_lrp_order_predict_response,omitempty"`
}

type AlibabaWdkScmLrpOrderPredictResponse struct {

    // result
    Result   *AlibabaWdkScmLrpOrderPredictApiResult `json:"result,omitempty"`

}
