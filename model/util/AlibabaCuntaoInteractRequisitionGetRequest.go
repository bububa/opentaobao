package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商获取物料申请单列表 APIRequest
alibaba.cuntao.interact.requisition.get

供应商获取物料申请单列表
*/
type AlibabaCuntaoInteractRequisitionGetRequest struct {
    model.Params

    // 页大小，默认20
    pageSize   int64 

    // 截止时间戳，开区间
    gmtCreateEnd   int64 

    // 开始时间戳，闭区间
    gmtCreateStart   int64 

    // 页码，从0开始
    pageIndex   int64 

}

func NewAlibabaCuntaoInteractRequisitionGetRequest() *AlibabaCuntaoInteractRequisitionGetRequest{
    return &AlibabaCuntaoInteractRequisitionGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetApiMethodName() string {
    return "alibaba.cuntao.interact.requisition.get"
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetGmtCreateEnd(gmtCreateEnd int64) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetGmtCreateEnd() int64 {
    return r.gmtCreateEnd
}

func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetGmtCreateStart(gmtCreateStart int64) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetGmtCreateStart() int64 {
    return r.gmtCreateStart
}

func (r *AlibabaCuntaoInteractRequisitionGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AlibabaCuntaoInteractRequisitionGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}

