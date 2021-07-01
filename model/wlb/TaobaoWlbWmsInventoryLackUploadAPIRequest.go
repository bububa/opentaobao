package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsInventoryLackUploadAPIRequest
缺货通知 API请求
taobao.wlb.wms.inventory.lack.upload

缺货通知 */
type TaobaoWlbWmsInventoryLackUploadAPIRequest struct {
	model.Params
	// 缺货通知信息
	_content *WlbWmsInventoryLackUpload
}

// New
