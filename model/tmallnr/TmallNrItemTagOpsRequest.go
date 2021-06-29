package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售商品打标去标 API请求
tmall.nr.item.tag.ops

参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
*/
type TmallNrItemTagOpsRequest struct {
    model.Params
    // 请求入参
    _tagReqDTO   *TagReqDto
}

// 初始化TmallNrItemTagOpsRequest对象
func NewTmallNrItemTagOpsRequest() *TmallNrItemTagOpsRequest{
    return &TmallNrItemTagOpsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrItemTagOpsRequest) GetApiMethodName() string {
    return "tmall.nr.item.tag.ops"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrItemTagOpsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagReqDTO Setter
// 请求入参
func (r *TmallNrItemTagOpsRequest) SetTagReqDTO(_tagReqDTO *TagReqDto) error {
    r._tagReqDTO = _tagReqDTO
    r.Set("tag_req_d_t_o", _tagReqDTO)
    return nil
}

// TagReqDTO Getter
func (r TmallNrItemTagOpsRequest) GetTagReqDTO() *TagReqDto {
    return r._tagReqDTO
}
