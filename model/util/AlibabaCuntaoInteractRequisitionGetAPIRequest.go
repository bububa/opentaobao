package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCuntaoInteractRequisitionGetAPIRequest 供应商获取物料申请单列表 API请求
// alibaba.cuntao.interact.requisition.get
//
// 供应商获取物料申请单列表
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

// NewAlibabaCuntaoInteractRequisitionGetRequest 初始化AlibabaCuntaoInteractRequisitionGetAPIRequest对象
func NewAlibabaCuntaoInteractRequisitionGetRequest() *AlibabaCuntaoInteractRequisitionGetAPIRequest {
	return &AlibabaCuntaoInteractRequisitionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetApiMethodName() string {
	return "alibaba.cuntao.interact.requisition.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageSize Setter
// 页大小，默认20
func (r *AlibabaCuntaoInteractRequisitionGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is GmtCreateEnd Setter
// 截止时间戳，开区间
func (r *AlibabaCuntaoInteractRequisitionGetAPIRequest) SetGmtCreateEnd(_gmtCreateEnd int64) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// Get GmtCreateEnd Getter
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetGmtCreateEnd() int64 {
	return r._gmtCreateEnd
}

// Set is GmtCreateStart Setter
// 开始时间戳，闭区间
func (r *AlibabaCuntaoInteractRequisitionGetAPIRequest) SetGmtCreateStart(_gmtCreateStart int64) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// Get GmtCreateStart Getter
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetGmtCreateStart() int64 {
	return r._gmtCreateStart
}

// Set is PageIndex Setter
// 页码，从0开始
func (r *AlibabaCuntaoInteractRequisitionGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r AlibabaCuntaoInteractRequisitionGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
