package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息同 API请求
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息
*/
type TmallNrtEasyhomememberSynRequest struct {
    model.Params
    // 入参
    _param   *ExternalMemberDTO
}

// 初始化TmallNrtEasyhomememberSynRequest对象
func NewTmallNrtEasyhomememberSynRequest() *TmallNrtEasyhomememberSynRequest{
    return &TmallNrtEasyhomememberSynRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtEasyhomememberSynRequest) GetApiMethodName() string {
    return "tmall.nrt.easyhomemember.syn"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtEasyhomememberSynRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *TmallNrtEasyhomememberSynRequest) SetParam(_param *ExternalMemberDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TmallNrtEasyhomememberSynRequest) GetParam() *ExternalMemberDTO {
    return r._param
}
