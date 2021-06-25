package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车租赁核销 APIRequest
tmall.car.lease.consume

租赁公司回传信息，核销
*/
type TmallCarLeaseConsumeRequest struct {
    model.Params

    // 核销请求
    cosumeCodeReqDTO   *CosumeCodeReqDto 

}

func NewTmallCarLeaseConsumeRequest() *TmallCarLeaseConsumeRequest{
    return &TmallCarLeaseConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeaseConsumeRequest) GetApiMethodName() string {
    return "tmall.car.lease.consume"
}

func (r TmallCarLeaseConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeaseConsumeRequest) SetCosumeCodeReqDTO(cosumeCodeReqDTO *CosumeCodeReqDto) error {
    r.cosumeCodeReqDTO = cosumeCodeReqDTO
    r.Set("cosume_code_req_d_t_o", cosumeCodeReqDTO)
    return nil
}

func (r TmallCarLeaseConsumeRequest) GetCosumeCodeReqDTO() *CosumeCodeReqDto {
    return r.cosumeCodeReqDTO
}

