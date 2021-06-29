package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
录入参加送彩票商品信息 API请求
taobao.caipiao.goods.info.input

录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
*/
type TaobaoCaipiaoGoodsInfoInputRequest struct {
    model.Params
    // 商品在淘宝的唯一id，不可为空
    goodsId   int64
    // 商品标题
    goodsTitle   string
    // 商品价格,保留两位小数，不可为空
    goodsPrice   float64
    // 商品主图地址
    goodsImage   string
    // 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
    presentType   int64
    // 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    actStartDate   string
    // 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    actEndDate   string
    // 商品类目编号，不可为空
    goodsType   int64
    // 彩种id,不可为空
    lotteryTypeId   int64
    // 店铺相关商品参加的送彩票活动描述
    goodsDesc   string
}

// 初始化TaobaoCaipiaoGoodsInfoInputRequest对象
func NewTaobaoCaipiaoGoodsInfoInputRequest() *TaobaoCaipiaoGoodsInfoInputRequest{
    return &TaobaoCaipiaoGoodsInfoInputRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetApiMethodName() string {
    return "taobao.caipiao.goods.info.input"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsId Setter
// 商品在淘宝的唯一id，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsId(goodsId int64) error {
    r.goodsId = goodsId
    r.Set("goods_id", goodsId)
    return nil
}

// GoodsId Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsId() int64 {
    return r.goodsId
}
// GoodsTitle Setter
// 商品标题
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsTitle(goodsTitle string) error {
    r.goodsTitle = goodsTitle
    r.Set("goods_title", goodsTitle)
    return nil
}

// GoodsTitle Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsTitle() string {
    return r.goodsTitle
}
// GoodsPrice Setter
// 商品价格,保留两位小数，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsPrice(goodsPrice float64) error {
    r.goodsPrice = goodsPrice
    r.Set("goods_price", goodsPrice)
    return nil
}

// GoodsPrice Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsPrice() float64 {
    return r.goodsPrice
}
// GoodsImage Setter
// 商品主图地址
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsImage(goodsImage string) error {
    r.goodsImage = goodsImage
    r.Set("goods_image", goodsImage)
    return nil
}

// GoodsImage Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsImage() string {
    return r.goodsImage
}
// PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetPresentType(presentType int64) error {
    r.presentType = presentType
    r.Set("present_type", presentType)
    return nil
}

// PresentType Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetPresentType() int64 {
    return r.presentType
}
// ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetActStartDate(actStartDate string) error {
    r.actStartDate = actStartDate
    r.Set("act_start_date", actStartDate)
    return nil
}

// ActStartDate Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetActStartDate() string {
    return r.actStartDate
}
// ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetActEndDate(actEndDate string) error {
    r.actEndDate = actEndDate
    r.Set("act_end_date", actEndDate)
    return nil
}

// ActEndDate Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetActEndDate() string {
    return r.actEndDate
}
// GoodsType Setter
// 商品类目编号，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsType(goodsType int64) error {
    r.goodsType = goodsType
    r.Set("goods_type", goodsType)
    return nil
}

// GoodsType Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsType() int64 {
    return r.goodsType
}
// LotteryTypeId Setter
// 彩种id,不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetLotteryTypeId(lotteryTypeId int64) error {
    r.lotteryTypeId = lotteryTypeId
    r.Set("lottery_type_id", lotteryTypeId)
    return nil
}

// LotteryTypeId Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetLotteryTypeId() int64 {
    return r.lotteryTypeId
}
// GoodsDesc Setter
// 店铺相关商品参加的送彩票活动描述
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsDesc(goodsDesc string) error {
    r.goodsDesc = goodsDesc
    r.Set("goods_desc", goodsDesc)
    return nil
}

// GoodsDesc Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsDesc() string {
    return r.goodsDesc
}
