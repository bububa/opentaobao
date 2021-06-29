package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新店铺基本信息 API请求
taobao.shop.update

目前只支持标题、公告和描述的更新
*/
type TaobaoShopUpdateRequest struct {
    model.Params
    // 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
    _title   string
    // 店铺公告。不超过1024个字符
    _bulletin   string
    // 店铺描述。10～2000个字符以内
    _desc   string
}

// 初始化TaobaoShopUpdateRequest对象
func NewTaobaoShopUpdateRequest() *TaobaoShopUpdateRequest{
    return &TaobaoShopUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoShopUpdateRequest) GetApiMethodName() string {
    return "taobao.shop.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoShopUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Title Setter
// 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
func (r *TaobaoShopUpdateRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoShopUpdateRequest) GetTitle() string {
    return r._title
}
// Bulletin Setter
// 店铺公告。不超过1024个字符
func (r *TaobaoShopUpdateRequest) SetBulletin(_bulletin string) error {
    r._bulletin = _bulletin
    r.Set("bulletin", _bulletin)
    return nil
}

// Bulletin Getter
func (r TaobaoShopUpdateRequest) GetBulletin() string {
    return r._bulletin
}
// Desc Setter
// 店铺描述。10～2000个字符以内
func (r *TaobaoShopUpdateRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoShopUpdateRequest) GetDesc() string {
    return r._desc
}
