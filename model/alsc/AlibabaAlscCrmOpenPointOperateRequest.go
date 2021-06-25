package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分操作接口 APIRequest
alibaba.alsc.crm.open.point.operate

同步积分接口
*/
type AlibabaAlscCrmOpenPointOperateRequest struct {
    model.Params

    // 入参
    paramPointOperateOpenReq   *PointOperateOpenReq 

}

func NewAlibabaAlscCrmOpenPointOperateRequest() *AlibabaAlscCrmOpenPointOperateRequest{
    return &AlibabaAlscCrmOpenPointOperateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenPointOperateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.point.operate"
}

func (r AlibabaAlscCrmOpenPointOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenPointOperateRequest) SetParamPointOperateOpenReq(paramPointOperateOpenReq *PointOperateOpenReq) error {
    r.paramPointOperateOpenReq = paramPointOperateOpenReq
    r.Set("param_point_operate_open_req", paramPointOperateOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenPointOperateRequest) GetParamPointOperateOpenReq() *PointOperateOpenReq {
    return r.paramPointOperateOpenReq
}

