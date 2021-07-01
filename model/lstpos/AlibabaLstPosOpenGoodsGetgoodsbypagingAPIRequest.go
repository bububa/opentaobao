package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest
分页查询用户全量的门店域商品接口(每页最多20条) API请求
alibaba.lst.pos.open.goods.getgoodsbypaging

分页查询用户全量的门店域商品接口(每页最多20条) */
type AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest struct {
	model.Params
	// 当前页码
	_page int64
	// 当前主账号userId
	_userId int64
}

// New
