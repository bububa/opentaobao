package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签接口 API请求
tmall.promotag.tag.find

查询用户创建的所有标签
*/
type TmallPromotagTagFindRequest struct {
    model.Params
    // 当前页码
    pageNo   int64
    // 每页显示个数
    pageSize   int64
    // 标签名称，查询时可选项
    tagName   string
    // 标签ID
    tagId   int64
}

// 初始化TmallPromotagTagFindRequest对象
func NewTmallPromotagTagFindRequest() *TmallPromotagTagFindRequest{
    return &TmallPromotagTagFindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTagFindRequest) GetApiMethodName() string {
    return "tmall.promotag.tag.find"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTagFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 当前页码
func (r *TmallPromotagTagFindRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TmallPromotagTagFindRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页显示个数
func (r *TmallPromotagTagFindRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallPromotagTagFindRequest) GetPageSize() int64 {
    return r.pageSize
}
// TagName Setter
// 标签名称，查询时可选项
func (r *TmallPromotagTagFindRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r TmallPromotagTagFindRequest) GetTagName() string {
    return r.tagName
}
// TagId Setter
// 标签ID
func (r *TmallPromotagTagFindRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTagFindRequest) GetTagId() int64 {
    return r.tagId
}
