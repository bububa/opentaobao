package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新店铺基本信息 APIRequest
taobao.shop.update

目前只支持标题、公告和描述的更新
*/
type TaobaoShopUpdateRequest struct {
    model.Params

    // 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个
    title   string 

    // 店铺公告。不超过1024个字符
    bulletin   string 

    // 店铺描述。10～2000个字符以内
    desc   string 

}

func NewTaobaoShopUpdateRequest() *TaobaoShopUpdateRequest{
    return &TaobaoShopUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoShopUpdateRequest) GetApiMethodName() string {
    return "taobao.shop.update"
}

func (r TaobaoShopUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoShopUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoShopUpdateRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoShopUpdateRequest) SetBulletin(bulletin string) error {
    r.bulletin = bulletin
    r.Set("bulletin", bulletin)
    return nil
}

func (r TaobaoShopUpdateRequest) GetBulletin() string {
    return r.bulletin
}

func (r *TaobaoShopUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoShopUpdateRequest) GetDesc() string {
    return r.desc
}

