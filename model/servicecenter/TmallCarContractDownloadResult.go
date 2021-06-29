package servicecenter
import (
    "github.com/bububa/opentaobao/model"
)

// TmallCarContractDownloadResult 
type TmallCarContractDownloadResult struct {
    // 当前时间
    GmtCurrentTime   int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
    // 错误码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回的合同二进制文件
    Objects   []*model.File `json:"objects,omitempty" xml:"objects>*model.File,omitempty"`
    // 成功与否
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 消耗时间
    CostTime   int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
}
