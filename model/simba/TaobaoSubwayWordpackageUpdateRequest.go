package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新词包 APIRequest
taobao.subway.wordpackage.update

批量更新词包
*/
type TaobaoSubwayWordpackageUpdateRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

    // 词包列表
    wordPackageDTOS   []ItemWordPackageDto 

}

func NewTaobaoSubwayWordpackageUpdateRequest() *TaobaoSubwayWordpackageUpdateRequest{
    return &TaobaoSubwayWordpackageUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubwayWordpackageUpdateRequest) GetApiMethodName() string {
    return "taobao.subway.wordpackage.update"
}

func (r TaobaoSubwayWordpackageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubwayWordpackageUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSubwayWordpackageUpdateRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSubwayWordpackageUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSubwayWordpackageUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSubwayWordpackageUpdateRequest) SetWordPackageDTOS(wordPackageDTOS []ItemWordPackageDto) error {
    r.wordPackageDTOS = wordPackageDTOS
    r.Set("word_package_d_t_o_s", wordPackageDTOS)
    return nil
}

func (r TaobaoSubwayWordpackageUpdateRequest) GetWordPackageDTOS() []ItemWordPackageDto {
    return r.wordPackageDTOS
}

