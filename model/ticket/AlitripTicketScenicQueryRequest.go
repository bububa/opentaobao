package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） API请求
alitrip.ticket.scenic.query

查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
*/
type AlitripTicketScenicQueryRequest struct {
    model.Params
    // 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
    _aliScenicId   int64
    // 当前分页。每页默认最多返回20条数据
    _currentPage   int64
    // 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
    _outScenicId   string
}

// 初始化AlitripTicketScenicQueryRequest对象
func NewAlitripTicketScenicQueryRequest() *AlitripTicketScenicQueryRequest{
    return &AlitripTicketScenicQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTicketScenicQueryRequest) GetApiMethodName() string {
    return "alitrip.ticket.scenic.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTicketScenicQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliScenicId Setter
// 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
func (r *AlitripTicketScenicQueryRequest) SetAliScenicId(_aliScenicId int64) error {
    r._aliScenicId = _aliScenicId
    r.Set("ali_scenic_id", _aliScenicId)
    return nil
}

// AliScenicId Getter
func (r AlitripTicketScenicQueryRequest) GetAliScenicId() int64 {
    return r._aliScenicId
}
// CurrentPage Setter
// 当前分页。每页默认最多返回20条数据
func (r *AlitripTicketScenicQueryRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTicketScenicQueryRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// OutScenicId Setter
// 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
func (r *AlitripTicketScenicQueryRequest) SetOutScenicId(_outScenicId string) error {
    r._outScenicId = _outScenicId
    r.Set("out_scenic_id", _outScenicId)
    return nil
}

// OutScenicId Getter
func (r AlitripTicketScenicQueryRequest) GetOutScenicId() string {
    return r._outScenicId
}
