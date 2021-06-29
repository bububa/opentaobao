package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询对账文件地址接口 API返回值 
taobao.elife.lifecard.recon

查询对账文件地址接口
*/
type TaobaoElifeLifecardReconAPIResponse struct {
    model.CommonResponse
    TaobaoElifeLifecardReconResponse
}

// 查询对账文件地址接口 成功返回结果
type TaobaoElifeLifecardReconResponse struct {
    XMLName xml.Name `xml:"elife_lifecard_recon_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 对账日期
    OpDate   string `json:"op_date,omitempty" xml:"op_date,omitempty"`
    // 成功标志
    Successed   bool `json:"successed,omitempty" xml:"successed,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 下载下载文件
    ReconFileUrl   string `json:"recon_file_url,omitempty" xml:"recon_file_url,omitempty"`
}
