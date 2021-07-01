package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslIteminfoBatchPutAPIRequest
批量写入商品信息接口 API请求
taobao.uscesl.iteminfo.batch.put

电子架签批量写入商品数据，用于电子价签展示 */
type TaobaoUsceslIteminfoBatchPutAPIRequest struct {
	model.Params
	// 商品变更信息集合
	_itemChangeBOList []ItemChangeBo
	// 门店ID
	_shopId int64
}

// New
