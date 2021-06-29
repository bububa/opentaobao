package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商获取物料申请单列表 API请求
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表
*/
type AlibabaCuntaoInteractRequisitionGetRequest struct {
    model.Params
    // 页大小，默认20
    _pageSize   int64
    // 截止时间戳，开区间
    _gmtCreateEnd   int64
    // 开始时间戳，闭区间
    _gmtCreateStart   int64
    // 页码，从0开始
    _pageIndex   int64
}

// 初始化AlibabaCuntaoInteractRequisitionGetRequest对象
func NewAlibabaCuntaoInteractRequisitionGetRequest() *AlibabaCuntaoInteractRequisitionGetRequest{
    return &AlibabaCuntaoInteractRequisitionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetApiMethodName() string {
    return "alibaba.cuntao.interact.requisition.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页大小，默认20
func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// GmtCreateEnd Setter
// 截止时间戳，开区间
func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetGmtCreateEnd(_gmtCreateEnd int64) error {
    r._gmtCreateEnd = _gmtCreateEnd
    r.Set("gmt_create_end", _gmtCreateEnd)
    return nil
}

// GmtCreateEnd Getter
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetGmtCreateEnd() int64 {
    return r._gmtCreateEnd
}
// GmtCreateStart Setter
// 开始时间戳，闭区间
func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetGmtCreateStart(_gmtCreateStart int64) error {
    r._gmtCreateStart = _gmtCreateStart
    r.Set("gmt_create_start", _gmtCreateStart)
    return nil
}

// GmtCreateStart Getter
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetGmtCreateStart() int64 {
    return r._gmtCreateStart
}
// PageIndex Setter
// 页码，从0开始
func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AlibabaCuntaoInteractRequisitionGetRequest) GetPageIndex() int64 {
    return r._pageIndex
}
