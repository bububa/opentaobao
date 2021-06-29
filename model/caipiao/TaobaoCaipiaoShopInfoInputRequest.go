package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
录入参加送彩票店铺信息 API请求
taobao.caipiao.shop.info.input

录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
*/
type TaobaoCaipiaoShopInfoInputRequest struct {
    model.Params
    // 店铺名称
    shopName   string
    // 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
    presentType   int64
    // 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    actStartDate   string
    // 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    actEndDate   string
    // 店铺类目编号，不可为空
    shopType   int64
    // 店铺参加的送彩票活动描述
    shopDesc   string
}

// 初始化TaobaoCaipiaoShopInfoInputRequest对象
func NewTaobaoCaipiaoShopInfoInputRequest() *TaobaoCaipiaoShopInfoInputRequest{
    return &TaobaoCaipiaoShopInfoInputRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoShopInfoInputRequest) GetApiMethodName() string {
    return "taobao.caipiao.shop.info.input"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoShopInfoInputRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopName Setter
// 店铺名称
func (r *TaobaoCaipiaoShopInfoInputRequest) SetShopName(shopName string) error {
    r.shopName = shopName
    r.Set("shop_name", shopName)
    return nil
}

// ShopName Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetShopName() string {
    return r.shopName
}
// PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaoCaipiaoShopInfoInputRequest) SetPresentType(presentType int64) error {
    r.presentType = presentType
    r.Set("present_type", presentType)
    return nil
}

// PresentType Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetPresentType() int64 {
    return r.presentType
}
// ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoShopInfoInputRequest) SetActStartDate(actStartDate string) error {
    r.actStartDate = actStartDate
    r.Set("act_start_date", actStartDate)
    return nil
}

// ActStartDate Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetActStartDate() string {
    return r.actStartDate
}
// ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoShopInfoInputRequest) SetActEndDate(actEndDate string) error {
    r.actEndDate = actEndDate
    r.Set("act_end_date", actEndDate)
    return nil
}

// ActEndDate Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetActEndDate() string {
    return r.actEndDate
}
// ShopType Setter
// 店铺类目编号，不可为空
func (r *TaobaoCaipiaoShopInfoInputRequest) SetShopType(shopType int64) error {
    r.shopType = shopType
    r.Set("shop_type", shopType)
    return nil
}

// ShopType Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetShopType() int64 {
    return r.shopType
}
// ShopDesc Setter
// 店铺参加的送彩票活动描述
func (r *TaobaoCaipiaoShopInfoInputRequest) SetShopDesc(shopDesc string) error {
    r.shopDesc = shopDesc
    r.Set("shop_desc", shopDesc)
    return nil
}

// ShopDesc Getter
func (r TaobaoCaipiaoShopInfoInputRequest) GetShopDesc() string {
    return r.shopDesc
}
