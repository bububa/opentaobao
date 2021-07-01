package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoGoodsInfoInputAPIRequest
录入参加送彩票商品信息 API请求
taobao.caipiao.goods.info.input

录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。 */
type TaobaoCaipiaoGoodsInfoInputAPIRequest struct {
	model.Params
	// 商品在淘宝的唯一id，不可为空
	_goodsId int64
	// 商品标题
	_goodsTitle string
	// 商品价格,保留两位小数，不可为空
	_goodsPrice float64
	// 商品主图地址
	_goodsImage string
	// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
	_presentType int64
	// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actStartDate string
	// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actEndDate string
	// 商品类目编号，不可为空
	_goodsType int64
	// 彩种id,不可为空
	_lotteryTypeId int64
	// 店铺相关商品参加的送彩票活动描述
	_goodsDesc string
}

// New
