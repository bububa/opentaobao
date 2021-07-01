package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品下载记录 API请求
taobao.fenxiao.distributor.items.get

供应商查询分销商商品下载记录。
*/
type TaobaoFenxiaoDistributorItemsGetAPIRequest struct {
    model.Params
    // 分销商ID 。
    _distributorId   int64
    // 设置开始时间。空为不设置。
    _startModified   string
    // 设置结束时间,空为不设置。
    _endModified   string
    // 页码（大于0的整数，默认1）
    _pageNo   int64
    // 每页记录数（默认20，最大50）
    _pageSize   int64
    // 产品ID
    _productId   int64
}

// 初始化TaobaoFenxiaoDistributorItemsGetAPIRequest对象
func NewTaobaoFenxiaoDistributorItemsGetRequest() *TaobaoFenxiaoDistributorItemsGetAPIRequest{
    return &TaobaoFenxiaoDistributorItemsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributor.items.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DistributorId Setter
// 分销商ID 。
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetDistributorId(_distributorId int64) error {
    r._distributorId = _distributorId
    r.Set("distributor_id", _distributorId)
    return nil
}

// DistributorId Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetDistributorId() int64 {
    return r._distributorId
}
// StartModified Setter
// 设置开始时间。空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 设置结束时间,空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetEndModified() string {
    return r._endModified
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoDistributorItemsGetAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoDistributorItemsGetAPIRequest) GetProductId() int64 {
    return r._productId
}
