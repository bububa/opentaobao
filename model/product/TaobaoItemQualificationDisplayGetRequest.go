package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资质采集配置异步获取接口 API请求
taobao.item.qualification.display.get

根据类目，商品，属性等参与动态获得资质采集配置
*/
type TaobaoItemQualificationDisplayGetRequest struct {
    model.Params
    // 参数列表，为key和value都是string的map的转化的json格式
    param   string
    // 商品id
    itemId   int64
    // 类目id
    categoryId   int64
}

// 初始化TaobaoItemQualificationDisplayGetRequest对象
func NewTaobaoItemQualificationDisplayGetRequest() *TaobaoItemQualificationDisplayGetRequest{
    return &TaobaoItemQualificationDisplayGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemQualificationDisplayGetRequest) GetApiMethodName() string {
    return "taobao.item.qualification.display.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemQualificationDisplayGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 参数列表，为key和value都是string的map的转化的json格式
func (r *TaobaoItemQualificationDisplayGetRequest) SetParam(param string) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r TaobaoItemQualificationDisplayGetRequest) GetParam() string {
    return r.param
}
// ItemId Setter
// 商品id
func (r *TaobaoItemQualificationDisplayGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoItemQualificationDisplayGetRequest) GetItemId() int64 {
    return r.itemId
}
// CategoryId Setter
// 类目id
func (r *TaobaoItemQualificationDisplayGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

// CategoryId Getter
func (r TaobaoItemQualificationDisplayGetRequest) GetCategoryId() int64 {
    return r.categoryId
}
