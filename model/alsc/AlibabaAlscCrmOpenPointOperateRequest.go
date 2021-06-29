package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分操作接口 API请求
alibaba.alsc.crm.open.point.operate

同步积分接口
*/
type AlibabaAlscCrmOpenPointOperateRequest struct {
    model.Params
    // 入参
    paramPointOperateOpenReq   *PointOperateOpenReq
}

// 初始化AlibabaAlscCrmOpenPointOperateRequest对象
func NewAlibabaAlscCrmOpenPointOperateRequest() *AlibabaAlscCrmOpenPointOperateRequest{
    return &AlibabaAlscCrmOpenPointOperateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenPointOperateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.point.operate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenPointOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPointOperateOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenPointOperateRequest) SetParamPointOperateOpenReq(paramPointOperateOpenReq *PointOperateOpenReq) error {
    r.paramPointOperateOpenReq = paramPointOperateOpenReq
    r.Set("param_point_operate_open_req", paramPointOperateOpenReq)
    return nil
}

// ParamPointOperateOpenReq Getter
func (r AlibabaAlscCrmOpenPointOperateRequest) GetParamPointOperateOpenReq() *PointOperateOpenReq {
    return r.paramPointOperateOpenReq
}
