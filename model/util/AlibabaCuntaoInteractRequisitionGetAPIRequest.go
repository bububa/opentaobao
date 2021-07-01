package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCuntaoInteractRequisitionGetAPIRequest
供应商获取物料申请单列表 API请求
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表 */
type AlibabaCuntaoInteractRequisitionGetAPIRequest struct {
	model.Params
	// 页大小，默认20
	_pageSize int64
	// 截止时间戳，开区间
	_gmtCreateEnd int64
	// 开始时间戳，闭区间
	_gmtCreateStart int64
	// 页码，从0开始
	_pageIndex int64
}

// New
