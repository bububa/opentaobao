package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广组 APIRequest
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组
*/
type TaobaoSimbaAdgroupsbyadgroupidsGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id列表
    adgroupIds   []Number 

    // 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
    pageSize   int64 

    // 页码，从1开始
    pageNo   int64 

}

func NewTaobaoSimbaAdgroupsbyadgroupidsGetRequest() *TaobaoSimbaAdgroupsbyadgroupidsGetRequest{
    return &TaobaoSimbaAdgroupsbyadgroupidsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupsbyadgroupids.get"
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetAdgroupIds(adgroupIds []Number) error {
    r.adgroupIds = adgroupIds
    r.Set("adgroup_ids", adgroupIds)
    return nil
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetAdgroupIds() []Number {
    return r.adgroupIds
}

func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

