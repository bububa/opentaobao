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
    _goodsId   int64
    // 商品标题
    _goodsTitle   string
    // 商品价格,保留两位小数，不可为空
    _goodsPrice   float64
    // 商品主图地址
    _goodsImage   string
    // 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
    _presentType   int64
    // 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    _actStartDate   string
    // 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
    _actEndDate   string
    // 商品类目编号，不可为空
    _goodsType   int64
    // 彩种id,不可为空
    _lotteryTypeId   int64
    // 店铺相关商品参加的送彩票活动描述
    _goodsDesc   string
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
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsId(_goodsId int64) error {
    r._goodsId = _goodsId
    r.Set("goods_id", _goodsId)
    return nil
}

// GoodsId Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsId() int64 {
    return r._goodsId
}
// GoodsTitle Setter
// 商品标题
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsTitle(_goodsTitle string) error {
    r._goodsTitle = _goodsTitle
    r.Set("goods_title", _goodsTitle)
    return nil
}

// GoodsTitle Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsTitle() string {
    return r._goodsTitle
}
// GoodsPrice Setter
// 商品价格,保留两位小数，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsPrice(_goodsPrice float64) error {
    r._goodsPrice = _goodsPrice
    r.Set("goods_price", _goodsPrice)
    return nil
}

// GoodsPrice Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsPrice() float64 {
    return r._goodsPrice
}
// GoodsImage Setter
// 商品主图地址
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsImage(_goodsImage string) error {
    r._goodsImage = _goodsImage
    r.Set("goods_image", _goodsImage)
    return nil
}

// GoodsImage Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsImage() string {
    return r._goodsImage
}
// PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetPresentType(_presentType int64) error {
    r._presentType = _presentType
    r.Set("present_type", _presentType)
    return nil
}

// PresentType Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetPresentType() int64 {
    return r._presentType
}
// ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetActStartDate(_actStartDate string) error {
    r._actStartDate = _actStartDate
    r.Set("act_start_date", _actStartDate)
    return nil
}

// ActStartDate Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetActStartDate() string {
    return r._actStartDate
}
// ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetActEndDate(_actEndDate string) error {
    r._actEndDate = _actEndDate
    r.Set("act_end_date", _actEndDate)
    return nil
}

// ActEndDate Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetActEndDate() string {
    return r._actEndDate
}
// GoodsType Setter
// 商品类目编号，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsType(_goodsType int64) error {
    r._goodsType = _goodsType
    r.Set("goods_type", _goodsType)
    return nil
}

// GoodsType Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsType() int64 {
    return r._goodsType
}
// LotteryTypeId Setter
// 彩种id,不可为空
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetLotteryTypeId(_lotteryTypeId int64) error {
    r._lotteryTypeId = _lotteryTypeId
    r.Set("lottery_type_id", _lotteryTypeId)
    return nil
}

// LotteryTypeId Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetLotteryTypeId() int64 {
    return r._lotteryTypeId
}
// GoodsDesc Setter
// 店铺相关商品参加的送彩票活动描述
func (r *TaobaoCaipiaoGoodsInfoInputRequest) SetGoodsDesc(_goodsDesc string) error {
    r._goodsDesc = _goodsDesc
    r.Set("goods_desc", _goodsDesc)
    return nil
}

// GoodsDesc Getter
func (r TaobaoCaipiaoGoodsInfoInputRequest) GetGoodsDesc() string {
    return r._goodsDesc
}
