package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新词包 API请求
taobao.subway.wordpackage.update

批量更新词包
*/
type TaobaoSubwayWordpackageUpdateRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组Id
    _adgroupId   int64
    // 词包列表
    _wordPackageDTOS   []ItemWordPackageDTO
}

// 初始化TaobaoSubwayWordpackageUpdateRequest对象
func NewTaobaoSubwayWordpackageUpdateRequest() *TaobaoSubwayWordpackageUpdateRequest{
    return &TaobaoSubwayWordpackageUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSubwayWordpackageUpdateRequest) GetApiMethodName() string {
    return "taobao.subway.wordpackage.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSubwayWordpackageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSubwayWordpackageUpdateRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSubwayWordpackageUpdateRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSubwayWordpackageUpdateRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSubwayWordpackageUpdateRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// WordPackageDTOS Setter
// 词包列表
func (r *TaobaoSubwayWordpackageUpdateRequest) SetWordPackageDTOS(_wordPackageDTOS []ItemWordPackageDTO) error {
    r._wordPackageDTOS = _wordPackageDTOS
    r.Set("word_package_d_t_o_s", _wordPackageDTOS)
    return nil
}

// WordPackageDTOS Getter
func (r TaobaoSubwayWordpackageUpdateRequest) GetWordPackageDTOS() []ItemWordPackageDTO {
    return r._wordPackageDTOS
}
