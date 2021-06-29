package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票品接口deleteItem APIRequest
alibaba.damai.mev.open.deleteitem

deleteItem
*/
type AlibabaDamaiMevOpenDeleteitemRequest struct {
    model.Params

    // 入参deleteItemParam
    deleteItemParam   *TicketItemIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeleteitemRequest() *AlibabaDamaiMevOpenDeleteitemRequest{
    return &AlibabaDamaiMevOpenDeleteitemRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeleteitemRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteitem"
}

func (r AlibabaDamaiMevOpenDeleteitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeleteitemRequest) SetDeleteItemParam(deleteItemParam *TicketItemIdOpenParam) error {
    r.deleteItemParam = deleteItemParam
    r.Set("delete_item_param", deleteItemParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeleteitemRequest) GetDeleteItemParam() *TicketItemIdOpenParam {
    return r.deleteItemParam
}

