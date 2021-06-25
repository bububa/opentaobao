package elife

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询对账文件地址接口 APIResponse
taobao.elife.lifecard.recon

查询对账文件地址接口
*/
type TaobaoElifeLifecardReconAPIResponse struct {
    model.CommonResponse
    Response *TaobaoElifeLifecardReconResponse `json:"taobao_elife_lifecard_recon_response,omitempty"`
}

type TaobaoElifeLifecardReconResponse struct {

    // 结果描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 对账日期
    OpDate   string `json:"op_date,omitempty"`

    // 成功标志
    Successed   bool `json:"successed,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 下载下载文件
    ReconFileUrl   string `json:"recon_file_url,omitempty"`

}
