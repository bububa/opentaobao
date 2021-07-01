package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCuntaoInteractRequisitionUpdateAPIRequest
更新物料制作状态 API请求
alibaba.cuntao.interact.requisition.update

村淘物料下沉，更新物料制作状态 */
type AlibabaCuntaoInteractRequisitionUpdateAPIRequest struct {
	model.Params
	// 物料制作状态
	_status string
	// 申请单id列表
	_uuidList []string
}

// New
