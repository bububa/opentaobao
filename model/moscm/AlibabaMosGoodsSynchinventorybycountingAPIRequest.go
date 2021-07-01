package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsSynchinventorybycountingAPIRequest
以盘点方式调整库存：传入商品实际库存 API请求
alibaba.mos.goods.synchinventorybycounting

以盘点方式调整库存：传入商品实际库存
盘点单自动判断数量增减 */
type AlibabaMosGoodsSynchinventorybycountingAPIRequest struct {
	model.Params
	// 专柜信息
	_paramCountingInfoDto *CountingInfoDto
	// 盘点库存项（最大列表长度：20）
	_countingItemDto []CountingItemDto
}

// New
