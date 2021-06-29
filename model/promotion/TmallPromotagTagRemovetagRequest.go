package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除标签定义 API请求
tmall.promotag.tag.removetag

用于删除标签定义，但是要确保目前该标签没有人群在使用。
*/
type TmallPromotagTagRemovetagRequest struct {
    model.Params
    // 需要删除的标签id
    _tagId   int64
}

// 初始化TmallPromotagTagRemovetagRequest对象
func NewTmallPromotagTagRemovetagRequest() *TmallPromotagTagRemovetagRequest{
    return &TmallPromotagTagRemovetagRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallPromotagTagRemovetagRequest) GetApiMethodName() string {
    return "tmall.promotag.tag.removetag"
}

// IRequest interface 方法, 获取API参数
func (r TmallPromotagTagRemovetagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagId Setter
// 需要删除的标签id
func (r *TmallPromotagTagRemovetagRequest) SetTagId(_tagId int64) error {
    r._tagId = _tagId
    r.Set("tag_id", _tagId)
    return nil
}

// TagId Getter
func (r TmallPromotagTagRemovetagRequest) GetTagId() int64 {
    return r._tagId
}
