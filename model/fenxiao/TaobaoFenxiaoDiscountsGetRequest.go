package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取折扣信息 API请求
taobao.fenxiao.discounts.get

查询折扣信息
*/
type TaobaoFenxiaoDiscountsGetRequest struct {
    model.Params
    // 折扣ID
    discountId   int64
    // 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
    extFields   string
}

// 初始化TaobaoFenxiaoDiscountsGetRequest对象
func NewTaobaoFenxiaoDiscountsGetRequest() *TaobaoFenxiaoDiscountsGetRequest{
    return &TaobaoFenxiaoDiscountsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDiscountsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.discounts.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDiscountsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DiscountId Setter
// 折扣ID
func (r *TaobaoFenxiaoDiscountsGetRequest) SetDiscountId(discountId int64) error {
    r.discountId = discountId
    r.Set("discount_id", discountId)
    return nil
}

// DiscountId Getter
func (r TaobaoFenxiaoDiscountsGetRequest) GetDiscountId() int64 {
    return r.discountId
}
// ExtFields Setter
// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
func (r *TaobaoFenxiaoDiscountsGetRequest) SetExtFields(extFields string) error {
    r.extFields = extFields
    r.Set("ext_fields", extFields)
    return nil
}

// ExtFields Getter
func (r TaobaoFenxiaoDiscountsGetRequest) GetExtFields() string {
    return r.extFields
}
