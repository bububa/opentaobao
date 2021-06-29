package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-标维度 API请求
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

// 初始化TaobaoQimenTagItemsQueryRequest对象
func NewTaobaoQimenTagItemsQueryRequest() *TaobaoQimenTagItemsQueryRequest{
    return &TaobaoQimenTagItemsQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTagItemsQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.tag.items.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTagItemsQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoQimenTagItemsQueryRequest) SetTagType(tagType string) error {
    r.tagType = tagType
    r.Set("tag_type", tagType)
    return nil
}

// TagType Getter
func (r TaobaoQimenTagItemsQueryRequest) GetTagType() string {
    return r.tagType
}
// Remark Setter
// 备注，string（500）
func (r *TaobaoQimenTagItemsQueryRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoQimenTagItemsQueryRequest) GetRemark() string {
    return r.remark
}
