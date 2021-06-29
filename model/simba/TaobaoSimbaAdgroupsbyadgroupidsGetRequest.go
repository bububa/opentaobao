package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广组 API请求
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组
*/
type TaobaoSimbaAdgroupsbyadgroupidsGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组Id列表
    _adgroupIds   []int64
    // 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
    _pageSize   int64
    // 页码，从1开始
    _pageNo   int64
}

// 初始化TaobaoSimbaAdgroupsbyadgroupidsGetRequest对象
func NewTaobaoSimbaAdgroupsbyadgroupidsGetRequest() *TaobaoSimbaAdgroupsbyadgroupidsGetRequest{
    return &TaobaoSimbaAdgroupsbyadgroupidsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupsbyadgroupids.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetNick() string {
    return r._nick
}
// AdgroupIds Setter
// 推广组Id列表
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetAdgroupIds(_adgroupIds []int64) error {
    r._adgroupIds = _adgroupIds
    r.Set("adgroup_ids", _adgroupIds)
    return nil
}

// AdgroupIds Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetAdgroupIds() []int64 {
    return r._adgroupIds
}
// PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 页码，从1开始
func (r *TaobaoSimbaAdgroupsbyadgroupidsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaAdgroupsbyadgroupidsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
