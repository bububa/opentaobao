package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车租赁核销 API请求
tmall.car.lease.consume

租赁公司回传信息，核销
*/
type TmallCarLeaseConsumeAPIRequest struct {
    model.Params
    // 核销请求
    _cosumeCodeReqDTO   *CosumeCodeReqDTO
}

// 初始化TmallCarLeaseConsumeAPIRequest对象
func NewTmallCarLeaseConsumeRequest() *TmallCarLeaseConsumeAPIRequest{
    return &TmallCarLeaseConsumeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseConsumeAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.consume"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseConsumeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CosumeCodeReqDTO Setter
// 核销请求
func (r *TmallCarLeaseConsumeAPIRequest) SetCosumeCodeReqDTO(_cosumeCodeReqDTO *CosumeCodeReqDTO) error {
    r._cosumeCodeReqDTO = _cosumeCodeReqDTO
    r.Set("cosume_code_req_d_t_o", _cosumeCodeReqDTO)
    return nil
}

// CosumeCodeReqDTO Getter
func (r TmallCarLeaseConsumeAPIRequest) GetCosumeCodeReqDTO() *CosumeCodeReqDTO {
    return r._cosumeCodeReqDTO
}
