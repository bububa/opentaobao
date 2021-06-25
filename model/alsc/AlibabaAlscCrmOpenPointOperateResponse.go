package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分操作接口 APIResponse
alibaba.alsc.crm.open.point.operate

同步积分接口
*/
type AlibabaAlscCrmOpenPointOperateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmOpenPointOperateResponse `json:"alibaba_alsc_crm_open_point_operate_response,omitempty"`
}

type AlibabaAlscCrmOpenPointOperateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
