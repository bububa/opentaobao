package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取本地模板信息 APIRequest
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板
*/
type TaobaoTanxNativetemplatesGetRequest struct {
    model.Params

    // dsp在tanx的memberid
    memberId   int64 

    // dsp对应的tanx的token
    token   string 

    // 1970年到现在的毫秒
    signTime   int64 

    // 本地模板ID列表
    templateIds   []int64 

}

func NewTaobaoTanxNativetemplatesGetRequest() *TaobaoTanxNativetemplatesGetRequest{
    return &TaobaoTanxNativetemplatesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxNativetemplatesGetRequest) GetApiMethodName() string {
    return "taobao.tanx.nativetemplates.get"
}

func (r TaobaoTanxNativetemplatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxNativetemplatesGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxNativetemplatesGetRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxNativetemplatesGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxNativetemplatesGetRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxNativetemplatesGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxNativetemplatesGetRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxNativetemplatesGetRequest) SetTemplateIds(templateIds []int64) error {
    r.templateIds = templateIds
    r.Set("template_ids", templateIds)
    return nil
}

func (r TaobaoTanxNativetemplatesGetRequest) GetTemplateIds() []int64 {
    return r.templateIds
}

