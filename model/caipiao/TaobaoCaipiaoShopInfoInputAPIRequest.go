package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoShopInfoInputAPIRequest
录入参加送彩票店铺信息 API请求
taobao.caipiao.shop.info.input

录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。 */
type TaobaoCaipiaoShopInfoInputAPIRequest struct {
	model.Params
	// 店铺名称
	_shopName string
	// 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空
	_presentType int64
	// 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actStartDate string
	// 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空
	_actEndDate string
	// 店铺类目编号，不可为空
	_shopType int64
	// 店铺参加的送彩票活动描述
	_shopDesc string
}

// New
