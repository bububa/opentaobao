package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据合同号查询铺位信息 APIResponse
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息
*/
type AlibabaMosBunkBunkinfoQuerybunkAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosBunkBunkinfoQuerybunkResponse `json:"alibaba_mos_bunk_bunkinfo_querybunk_response,omitempty"`
}

type AlibabaMosBunkBunkinfoQuerybunkResponse struct {

    // 返回结果
    Result   *AlibabaMosBunkBunkinfoQuerybunkResultDo `json:"result,omitempty"`

}
