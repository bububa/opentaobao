package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-出票 API请求
alibaba.damai.maitix.order.confirm

出票
*/
type AlibabaDamaiMaitixOrderConfirmAPIRequest struct {
    model.Params
    // 出票入参
    _param   *MoaConfirmOrderParam
}

// 初始化AlibabaDamaiMaitixOrderConfirmAPIRequest对象
func NewAlibabaDamaiMaitixOrderConfirmRequest() *AlibabaDamaiMaitixOrderConfirmAPIRequest{
    return &AlibabaDamaiMaitixOrderConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 出票入参
func (r *AlibabaDamaiMaitixOrderConfirmAPIRequest) SetParam(_param *MoaConfirmOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixOrderConfirmAPIRequest) GetParam() *MoaConfirmOrderParam {
    return r._param
}
