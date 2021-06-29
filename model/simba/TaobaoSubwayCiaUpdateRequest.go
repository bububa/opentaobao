package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改单元智能出价 API请求
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

// 初始化TaobaoSubwayCiaUpdateRequest对象
func NewTaobaoSubwayCiaUpdateRequest() *TaobaoSubwayCiaUpdateRequest{
    return &TaobaoSubwayCiaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaUpdateRequest) GetApiMethodName() string {
    return "taobao.subway.cia.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayCiaUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayCiaUpdateRequest) GetNick() string {
    return r.nick
}
// CiaConfigs Setter
// 系统自动生成
func (r *TaobaoSubwayCiaUpdateRequest) SetCiaConfigs(ciaConfigs []CiaUpdateDto) error {
    r.ciaConfigs = ciaConfigs
    r.Set("cia_configs", ciaConfigs)
    return nil
}

// CiaConfigs Getter
func (r TaobaoSubwayCiaUpdateRequest) GetCiaConfigs() []CiaUpdateDto {
    return r.ciaConfigs
}
