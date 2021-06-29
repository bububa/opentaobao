package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动详情列表 API请求
taobao.ump.details.get

分页查询优惠详情列表
*/
type TaobaoUmpDetailsGetRequest struct {
    model.Params
    // 营销活动id
    _actId   int64
    // 分页的页码
    _pageNo   int64
    // 每页的最大条数
    _pageSize   int64
}

// 初始化TaobaoUmpDetailsGetRequest对象
func NewTaobaoUmpDetailsGetRequest() *TaobaoUmpDetailsGetRequest{
    return &TaobaoUmpDetailsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailsGetRequest) GetApiMethodName() string {
    return "taobao.ump.details.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActId Setter
// 营销活动id
func (r *TaobaoUmpDetailsGetRequest) SetActId(_actId int64) error {
    r._actId = _actId
    r.Set("act_id", _actId)
    return nil
}

// ActId Getter
func (r TaobaoUmpDetailsGetRequest) GetActId() int64 {
    return r._actId
}
// PageNo Setter
// 分页的页码
func (r *TaobaoUmpDetailsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoUmpDetailsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页的最大条数
func (r *TaobaoUmpDetailsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoUmpDetailsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
