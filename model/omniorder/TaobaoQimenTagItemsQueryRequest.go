package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-标维度 APIRequest
taobao.qimen.tag.items.query

调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
*/
type TaobaoQimenTagItemsQueryRequest struct {
    model.Params

    // 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
    tagType   string 

    // 备注，string（500）
    remark   string 

}

func NewTaobaoQimenTagItemsQueryRequest() *TaobaoQimenTagItemsQueryRequest{
    return &TaobaoQimenTagItemsQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTagItemsQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.tag.items.query"
}

func (r TaobaoQimenTagItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTagItemsQueryRequest) SetTagType(tagType string) error {
    r.tagType = tagType
    r.Set("tag_type", tagType)
    return nil
}

func (r TaobaoQimenTagItemsQueryRequest) GetTagType() string {
    return r.tagType
}

func (r *TaobaoQimenTagItemsQueryRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TaobaoQimenTagItemsQueryRequest) GetRemark() string {
    return r.remark
}

