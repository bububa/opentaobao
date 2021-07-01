package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemstoreBandingAPIRequest
商品关联绑定接口 API请求
taobao.qimen.itemstore.banding

商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200 */
type TaobaoQimenItemstoreBandingAPIRequest struct {
	model.Params
	// 门店列表
	_storeIds []int64
	// 备注信息
	_remark string
	// 操作类型
	_actionType string
	// 线上商品ID
	_itemId int64
}

// New
