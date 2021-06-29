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
    distributorId   int64
    // 设置开始时间。空为不设置。
    startModified   string
    // 设置结束时间,空为不设置。
    endModified   string
    // 页码（大于0的整数，默认1）
    pageNo   int64
    // 每页记录数（默认20，最大50）
    pageSize   int64
    // 产品ID
    productId   int64
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
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetDistributorId(distributorId int64) error {
    r.distributorId = distributorId
    r.Set("distributor_id", distributorId)
    return nil
}

// DistributorId Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetDistributorId() int64 {
    return r.distributorId
}
// StartModified Setter
// 设置开始时间。空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

// StartModified Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetStartModified() string {
    return r.startModified
}
// EndModified Setter
// 设置结束时间,空为不设置。
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

// EndModified Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetEndModified() string {
    return r.endModified
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoDistributorItemsGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoDistributorItemsGetRequest) GetProductId() int64 {
    return r.productId
}
