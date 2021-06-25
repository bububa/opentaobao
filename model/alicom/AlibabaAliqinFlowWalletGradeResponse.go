package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取流量档位 APIResponse
alibaba.aliqin.flow.wallet.grade

获取直充流量档位
*/
type AlibabaAliqinFlowWalletGradeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowWalletGradeResponse `json:"alibaba_aliqin_flow_wallet_grade_response,omitempty"`
}

type AlibabaAliqinFlowWalletGradeResponse struct {

    // 档位
    Grade   string `json:"grade,omitempty"`

}
