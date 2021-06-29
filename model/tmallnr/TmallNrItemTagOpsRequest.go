package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售商品打标去标 APIRequest
tmall.nr.item.tag.ops

参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
*/
type TmallNrItemTagOpsRequest struct {
    model.Params

    // 请求入参
    tagReqDTO   *TagReqDto 

}

func NewTmallNrItemTagOpsRequest() *TmallNrItemTagOpsRequest{
    return &TmallNrItemTagOpsRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrItemTagOpsRequest) GetApiMethodName() string {
    return "tmall.nr.item.tag.ops"
}

func (r TmallNrItemTagOpsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrItemTagOpsRequest) SetTagReqDTO(tagReqDTO *TagReqDto) error {
    r.tagReqDTO = tagReqDTO
    r.Set("tag_req_d_t_o", tagReqDTO)
    return nil
}

func (r TmallNrItemTagOpsRequest) GetTagReqDTO() *TagReqDto {
    return r.tagReqDTO
}

