package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合作申请查询 API请求
taobao.fenxiao.requisitions.get

合作申请查询
*/
type TaobaoFenxiaoRequisitionsGetRequest struct {
    model.Params
    // 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
    status   int64
    // 申请开始时间yyyy-MM-dd
    applyStart   string
    // 申请结束时间yyyy-MM-dd
    applyEnd   string
    // 页码（大于0的整数，默认1）
    pageNo   int64
    // 每页记录数（默认20，最大50）
    pageSize   int64
}

// 初始化TaobaoFenxiaoRequisitionsGetRequest对象
func NewTaobaoFenxiaoRequisitionsGetRequest() *TaobaoFenxiaoRequisitionsGetRequest{
    return &TaobaoFenxiaoRequisitionsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoRequisitionsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.requisitions.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoRequisitionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
func (r *TaobaoFenxiaoRequisitionsGetRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoFenxiaoRequisitionsGetRequest) GetStatus() int64 {
    return r.status
}
// ApplyStart Setter
// 申请开始时间yyyy-MM-dd
func (r *TaobaoFenxiaoRequisitionsGetRequest) SetApplyStart(applyStart string) error {
    r.applyStart = applyStart
    r.Set("apply_start", applyStart)
    return nil
}

// ApplyStart Getter
func (r TaobaoFenxiaoRequisitionsGetRequest) GetApplyStart() string {
    return r.applyStart
}
// ApplyEnd Setter
// 申请结束时间yyyy-MM-dd
func (r *TaobaoFenxiaoRequisitionsGetRequest) SetApplyEnd(applyEnd string) error {
    r.applyEnd = applyEnd
    r.Set("apply_end", applyEnd)
    return nil
}

// ApplyEnd Getter
func (r TaobaoFenxiaoRequisitionsGetRequest) GetApplyEnd() string {
    return r.applyEnd
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoRequisitionsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoRequisitionsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoRequisitionsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoRequisitionsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
