package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取流量档位 APIResponse
alibaba.aliqin.flow.wallet.grade

获取直充流量档位
*/
type AlibabaAliqinFlowWalletGradeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_wallet_grade_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 档位
    
    Grade   string `json:"grade,omitempty" xml:"