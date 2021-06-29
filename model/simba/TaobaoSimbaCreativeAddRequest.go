package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加创意 API请求
taobao.simba.creative.add

创建一个创意
*/
type TaobaoSimbaCreativeAddRequest struct {
    model.Params
    // 推广组Id
    _adgroupId   int64
    // 创意标题，最多20个汉字
    _title   string
    // 创意图片地址，必须是推广组对应商品的图片之一
    _imgUrl   string
    // 主人昵称
    _nick   string
}

// 初始化TaobaoSimbaCreativeAddRequest对象
func NewTaobaoSimbaCreativeAddRequest() *TaobaoSimbaCreativeAddRequest{
    return &TaobaoSimbaCreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeAddRequest) GetApiMethodName() string {
    return "taobao.simba.creative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativeAddRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaCreativeAddRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaCreativeAddRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoSimbaCreativeAddRequest) GetTitle() string {
    return r._title
}
// ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaCreativeAddRequest) SetImgUrl(_imgUrl string) error {
    r._imgUrl = _imgUrl
    r.Set("img_url", _imgUrl)
    return nil
}

// ImgUrl Getter
func (r TaobaoSimbaCreativeAddRequest) GetImgUrl() string {
    return r._imgUrl
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeAddRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeAddRequest) GetNick() string {
    return r._nick
}
