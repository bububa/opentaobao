package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增量获取卖家会员 API请求
taobao.crm.members.increment.get

增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
*/
type TaobaoCrmMembersIncrementGetRequest struct {
    model.Params
    // 会员等级
    grade   int64
    // 卖家修改会员信息的时间起点.
    startModify   string
    // 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
    endModify   string
    // 每页显示的会员数，page_size的值不能超过100，最小值要大于1
    pageSize   int64
    // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
    currentPage   int64
}

// 初始化TaobaoCrmMembersIncrementGetRequest对象
func NewTaobaoCrmMembersIncrementGetRequest() *TaobaoCrmMembersIncrementGetRequest{
    return &TaobaoCrmMembersIncrementGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersIncrementGetRequest) GetApiMethodName() string {
    return "taobao.crm.members.increment.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersIncrementGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Grade Setter
// 会员等级
func (r *TaobaoCrmMembersIncrementGetRequest) SetGrade(grade int64) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

// Grade Getter
func (r TaobaoCrmMembersIncrementGetRequest) GetGrade() int64 {
    return r.grade
}
// StartModify Setter
// 卖家修改会员信息的时间起点.
func (r *TaobaoCrmMembersIncrementGetRequest) SetStartModify(startModify string) error {
    r.startModify = startModify
    r.Set("start_modify", startModify)
    return nil
}

// StartModify Getter
func (r TaobaoCrmMembersIncrementGetRequest) GetStartModify() string {
    return r.startModify
}
// EndModify Setter
// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
func (r *TaobaoCrmMembersIncrementGetRequest) SetEndModify(endModify string) error {
    r.endModify = endModify
    r.Set("end_modify", endModify)
    return nil
}

// EndModify Getter
func (r TaobaoCrmMembersIncrementGetRequest) GetEndModify() string {
    return r.endModify
}
// PageSize Setter
// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
func (r *TaobaoCrmMembersIncrementGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoCrmMembersIncrementGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
func (r *TaobaoCrmMembersIncrementGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoCrmMembersIncrementGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
