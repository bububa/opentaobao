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
type AlibabaDamaiMaitixOrderConfirmRequest struct {
    model.Params
    // 出票入参
    _param   *MoaConfirmOrderParam
}

// 初始化AlibabaDamaiMaitixOrderConfirmRequest对象
func NewAlibabaDamaiMaitixOrderConfirmRequest() *AlibabaDamaiMaitixOrderConfirmRequest{
    return &AlibabaDamaiMaitixOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderConfirmRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 出票入参
func (r *AlibabaDamaiMaitixOrderConfirmRequest) SetParam(_param *MoaConfirmOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixOrderConfirmRequest) GetParam() *MoaConfirmOrderParam {
    return r._param
}
