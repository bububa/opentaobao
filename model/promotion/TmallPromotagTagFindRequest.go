package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签接口 APIRequest
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

func NewTmallPromotagTagFindRequest() *TmallPromotagTagFindRequest{
    return &TmallPromotagTagFindRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPromotagTagFindRequest) GetApiMethodName() string {
    return "tmall.promotag.tag.find"
}

func (r TmallPromotagTagFindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPromotagTagFindRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TmallPromotagTagFindRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TmallPromotagTagFindRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallPromotagTagFindRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallPromotagTagFindRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

func (r TmallPromotagTagFindRequest) GetTagName() string {
    return r.tagName
}

func (r *TmallPromotagTagFindRequest) SetTagId(tagId int64) error {
    r.tagId = tagId
    r.Set("tag_id", tagId)
    return nil
}

func (r TmallPromotagTagFindRequest) GetTagId() int64 {
    return r.tagId
}

