package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoGoodsInfoInputAPIRequest 录入参加送彩票商品信息 API请求
// taobao.caipiao.goods.info.input
//
// 录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
type TaobaoCaipiaoGoodsInfoInputAPIRequest struct {
	model.Params
	// 商品标题
	_goodsTitle string
	// 商品主图地址
	_goodsImage string
	// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actStartDate string
	// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actEndDate string
	// 店铺相关商品参加的送彩票活动描述
	_goodsDesc string
	// 商品在淘宝的唯一id，不可为空
	_goodsId int64
	// 商品价格,保留两位小数，不可为空
	_goodsPrice float64
	// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
	_presentType int64
	// 商品类目编号，不可为空
	_goodsType int64
	// 彩种id,不可为空
	_lotteryTypeId int64
}

// NewTaobaoCaipiaoGoodsInfoInputRequest 初始化TaobaoCaipiaoGoodsInfoInputAPIRequest对象
func NewTaobaoCaipiaoGoodsInfoInputRequest() *TaobaoCaipiaoGoodsInfoInputAPIRequest {
	return &TaobaoCaipiaoGoodsInfoInputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.goods.info.input"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsTitle is GoodsTitle Setter
// 商品标题
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsTitle(_goodsTitle string) error {
	r._goodsTitle = _goodsTitle
	r.Set("goods_title", _goodsTitle)
	return nil
}

// GetGoodsTitle GoodsTitle Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsTitle() string {
	return r._goodsTitle
}

// SetGoodsImage is GoodsImage Setter
// 商品主图地址
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsImage(_goodsImage string) error {
	r._goodsImage = _goodsImage
	r.Set("goods_image", _goodsImage)
	return nil
}

// GetGoodsImage GoodsImage Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsImage() string {
	return r._goodsImage
}

// SetActStartDate is ActStartDate Setter
// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetActStartDate(_actStartDate string) error {
	r._actStartDate = _actStartDate
	r.Set("act_start_date", _actStartDate)
	return nil
}

// GetActStartDate ActStartDate Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetActStartDate() string {
	return r._actStartDate
}

// SetActEndDate is ActEndDate Setter
// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetActEndDate(_actEndDate string) error {
	r._actEndDate = _actEndDate
	r.Set("act_end_date", _actEndDate)
	return nil
}

// GetActEndDate ActEndDate Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetActEndDate() string {
	return r._actEndDate
}

// SetGoodsDesc is GoodsDesc Setter
// 店铺相关商品参加的送彩票活动描述
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsDesc(_goodsDesc string) error {
	r._goodsDesc = _goodsDesc
	r.Set("goods_desc", _goodsDesc)
	return nil
}

// GetGoodsDesc GoodsDesc Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsDesc() string {
	return r._goodsDesc
}

// SetGoodsId is GoodsId Setter
// 商品在淘宝的唯一id，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsId(_goodsId int64) error {
	r._goodsId = _goodsId
	r.Set("goods_id", _goodsId)
	return nil
}

// GetGoodsId GoodsId Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsId() int64 {
	return r._goodsId
}

// SetGoodsPrice is GoodsPrice Setter
// 商品价格,保留两位小数，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsPrice(_goodsPrice float64) error {
	r._goodsPrice = _goodsPrice
	r.Set("goods_price", _goodsPrice)
	return nil
}

// GetGoodsPrice GoodsPrice Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsPrice() float64 {
	return r._goodsPrice
}

// SetPresentType is PresentType Setter
// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetPresentType(_presentType int64) error {
	r._presentType = _presentType
	r.Set("present_type", _presentType)
	return nil
}

// GetPresentType PresentType Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetPresentType() int64 {
	return r._presentType
}

// SetGoodsType is GoodsType Setter
// 商品类目编号，不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetGoodsType(_goodsType int64) error {
	r._goodsType = _goodsType
	r.Set("goods_type", _goodsType)
	return nil
}

// GetGoodsType GoodsType Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetGoodsType() int64 {
	return r._goodsType
}

// SetLotteryTypeId is LotteryTypeId Setter
// 彩种id,不可为空
func (r *TaobaoCaipiaoGoodsInfoInputAPIRequest) SetLotteryTypeId(_lotteryTypeId int64) error {
	r._lotteryTypeId = _lotteryTypeId
	r.Set("lottery_type_id", _lotteryTypeId)
	return nil
}

// GetLotteryTypeId LotteryTypeId Getter
func (r TaobaoCaipiaoGoodsInfoInputAPIRequest) GetLotteryTypeId() int64 {
	return r._lotteryTypeId
}
