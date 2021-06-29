package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取本地模板信息 API请求
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板
*/
type TaobaoTanxNativetemplatesGetRequest struct {
    model.Params
    // dsp在tanx的memberid
    _memberId   int64
    // dsp对应的tanx的token
    _token   string
    // 1970年到现在的毫秒
    _signTime   int64
    // 本地模板ID列表
    _templateIds   []int64
}

// 初始化TaobaoTanxNativetemplatesGetRequest对象
func NewTaobaoTanxNativetemplatesGetRequest() *TaobaoTanxNativetemplatesGetRequest{
    return &TaobaoTanxNativetemplatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxNativetemplatesGetRequest) GetApiMethodName() string {
    return "taobao.tanx.nativetemplates.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxNativetemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// dsp在tanx的memberid
func (r *TaobaoTanxNativetemplatesGetRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxNativetemplatesGetRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp对应的tanx的token
func (r *TaobaoTanxNativetemplatesGetRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxNativetemplatesGetRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 1970年到现在的毫秒
func (r *TaobaoTanxNativetemplatesGetRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxNativetemplatesGetRequest) GetSignTime() int64 {
    return r._signTime
}
// TemplateIds Setter
// 本地模板ID列表
func (r *TaobaoTanxNativetemplatesGetRequest) SetTemplateIds(_templateIds []int64) error {
    r._templateIds = _templateIds
    r.Set("template_ids", _templateIds)
    return nil
}

// TemplateIds Getter
func (r TaobaoTanxNativetemplatesGetRequest) GetTemplateIds() []int64 {
    return r._templateIds
}
