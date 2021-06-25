package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改单元智能出价 APIRequest
taobao.subway.cia.update

批量修改直通车推广单元的智能出价配置
*/
type TaobaoSubwayCiaUpdateRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 系统自动生成
    ciaConfigs   []CiaUpdateDto 

}

func NewTaobaoSubwayCiaUpdateRequest() *TaobaoSubwayCiaUpdateRequest{
    return &TaobaoSubwayCiaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubwayCiaUpdateRequest) GetApiMethodName() string {
    return "taobao.subway.cia.update"
}

func (r TaobaoSubwayCiaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubwayCiaUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSubwayCiaUpdateRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSubwayCiaUpdateRequest) SetCiaConfigs(ciaConfigs []CiaUpdateDto) error {
    r.ciaConfigs = ciaConfigs
    r.Set("cia_configs", ciaConfigs)
    return nil
}

func (r TaobaoSubwayCiaUpdateRequest) GetCiaConfigs() []CiaUpdateDto {
    return r.ciaConfigs
}

