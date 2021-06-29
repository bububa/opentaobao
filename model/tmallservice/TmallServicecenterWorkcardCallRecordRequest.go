package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回访记录回传API API请求
tmall.servicecenter.workcard.call.record

客满回访信息登记
*/
type TmallServicecenterWorkcardCallRecordRequest struct {
    model.Params
    // 请求入参
    _busiRequest   *UpdateAttributeRequest
}

// 初始化TmallServicecenterWorkcardCallRecordRequest对象
func NewTmallServicecenterWorkcardCallRecordRequest() *TmallServicecenterWorkcardCallRecordRequest{
    return &TmallServicecenterWorkcardCallRecordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCallRecordRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.call.record"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCallRecordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusiRequest Setter
// 请求入参
func (r *TmallServicecenterWorkcardCallRecordRequest) SetBusiRequest(_busiRequest *UpdateAttributeRequest) error {
    r._busiRequest = _busiRequest
    r.Set("busi_request", _busiRequest)
    return nil
}

// BusiRequest Getter
func (r TmallServicecenterWorkcardCallRecordRequest) GetBusiRequest() *UpdateAttributeRequest {
    return r._busiRequest
}
