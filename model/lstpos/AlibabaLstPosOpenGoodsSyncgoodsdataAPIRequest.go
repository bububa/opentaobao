package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest
门店商品批量同步接口(最多10条商品信息) API请求
alibaba.lst.pos.open.goods.syncgoodsdata

门店商品批量同步接口(最多10条商品信息) */
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest struct {
	model.Params
	// 商品对象列表
	_goodsDTOList []GoodsDto
	// 用户主账号Id
	_userId int64
}

// New
