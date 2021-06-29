package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤 API请求
taobao.openim.snfilterword.setfilter

设置openim关键词过滤
*/
type TaobaoOpenimSnfilterwordSetfilterRequest struct {
    model.Params
    // 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
    _creator   string
    // 需要过滤的关键词
    _filterword   string
    // 过滤原因描述
    _desc   string
}

// 初始化TaobaoOpenimSnfilterwordSetfilterRequest对象
func NewTaobaoOpenimSnfilterwordSetfilterRequest() *TaobaoOpenimSnfilterwordSetfilterRequest{
    return &TaobaoOpenimSnfilterwordSetfilterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetApiMethodName() string {
    return "taobao.openim.snfilterword.setfilter"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Creator Setter
// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetCreator(_creator string) error {
    r._creator = _creator
    r.Set("creator", _creator)
    return nil
}

// Creator Getter
func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetCreator() string {
    return r._creator
}
// Filterword Setter
// 需要过滤的关键词
func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetFilterword(_filterword string) error {
    r._filterword = _filterword
    r.Set("filterword", _filterword)
    return nil
}

// Filterword Getter
func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetFilterword() string {
    return r._filterword
}
// Desc Setter
// 过滤原因描述
func (r *TaobaoOpenimSnfilterwordSetfilterRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoOpenimSnfilterwordSetfilterRequest) GetDesc() string {
    return r._desc
}
