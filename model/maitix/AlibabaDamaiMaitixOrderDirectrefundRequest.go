package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-直接退票 API请求
alibaba.damai.maitix.order.directrefund

大麦-退票
*/
type AlibabaDamaiMaitixOrderDirectrefundRequest struct {
    model.Params
    // 退票入参
    _param   *MoaRefundAuditParam
}

// 初始化AlibabaDamaiMaitixOrderDirectrefundRequest对象
func NewAlibabaDamaiMaitixOrderDirectrefundRequest() *AlibabaDamaiMaitixOrderDirectrefundRequest{
    return &AlibabaDamaiMaitixOrderDirectrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.directrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 退票入参
func (r *AlibabaDamaiMaitixOrderDirectrefundRequest) SetParam(_param *MoaRefundAuditParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixOrderDirectrefundRequest) GetParam() *MoaRefundAuditParam {
    return r._param
}
