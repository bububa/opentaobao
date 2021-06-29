package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后信息同步 APIRequest
tmall.car.lease.postsynchronize

商家同步天猫开新车租后方案
*/
type TmallCarLeasePostsynchronizeRequest struct {
    model.Params

    // 租后方案信息
    schemeDto   *CarLeasePostSchemeSynchronizeDto 

}

func NewTmallCarLeasePostsynchronizeRequest() *TmallCarLeasePostsynchronizeRequest{
    return &TmallCarLeasePostsynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallCarLeasePostsynchronizeRequest) GetApiMethodName() string {
    return "tmall.car.lease.postsynchronize"
}

func (r TmallCarLeasePostsynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallCarLeasePostsynchronizeRequest) SetSchemeDto(schemeDto *CarLeasePostSchemeSynchronizeDto) error {
    r.schemeDto = schemeDto
    r.Set("scheme_dto", schemeDto)
    return nil
}

func (r TmallCarLeasePostsynchronizeRequest) GetSchemeDto() *CarLeasePostSchemeSynchronizeDto {
    return r.schemeDto
}

