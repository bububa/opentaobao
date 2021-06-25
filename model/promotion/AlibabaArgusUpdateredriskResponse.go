package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新红线价格 APIResponse
alibaba.argus.updateredrisk

商品健康中心新增红线价格规则
*/
type AlibabaArgusUpdateredriskAPIResponse struct {
    model.CommonResponse
    Response *AlibabaArgusUpdateredriskResponse `json:"alibaba_argus_updateredrisk_response,omitempty"`
}

type AlibabaArgusUpdateredriskResponse struct {

    // 结果是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ReturnCode   string `json:"return_code,omitempty"`

    // 总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
