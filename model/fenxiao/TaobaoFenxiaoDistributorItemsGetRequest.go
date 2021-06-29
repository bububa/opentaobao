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
type TaobaoFenxiaoDistributorItemsGetRequest struct {
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

// 初始化TaobaoFenxiaoDistributorItemsGetRequest对象
func NewTaobaoFenxiaoDistributorItemsGetRequest() *TaobaoFenxiaoDistributorItemsGetRequest{
    return &TaobaoFenxiaoDistributorItemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.distributor.items.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DistributorId Setter
// 分销商ID 。
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetDistributorId(_distributorId int64) error {
    r._distributorId = _distributorId
    r.Set("distributor_id", _distributorId)
    return nil
}

// DistributorId Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetDistributorId() int64 {
    return r._distributorId
}
// StartModified Setter
// 设置开始时间。空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 设置结束时间,空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetEndModified() string {
    return r._endModified
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetProductId() int64 {
    return r._productId
}
