package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息同 APIRequest
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息
*/
type TmallNrtEasyhomememberSynRequest struct {
    model.Params

    // 入参
    param   *ExternalMemberDTO 

}

func NewTmallNrtEasyhomememberSynRequest() *TmallNrtEasyhomememberSynRequest{
    return &TmallNrtEasyhomememberSynRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtEasyhomememberSynRequest) GetApiMethodName() string {
    return "tmall.nrt.easyhomemember.syn"
}

func (r TmallNrtEasyhomememberSynRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtEasyhomememberSynRequest) SetParam(param *ExternalMemberDTO) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TmallNrtEasyhomememberSynRequest) GetParam() *ExternalMemberDTO {
    return r.param
}

