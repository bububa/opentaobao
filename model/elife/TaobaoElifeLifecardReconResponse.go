package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询对账文件地址接口 APIResponse
taobao.elife.lifecard.recon

查询对账文件地址接口
*/
type TaobaoElifeLifecardReconAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"elife_lifecard_recon_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"