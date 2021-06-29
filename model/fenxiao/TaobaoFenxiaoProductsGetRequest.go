package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询产品列表 API请求
taobao.fenxiao.products.get

查询供应商的产品数据。<br/><br/>    * 入参传入pids将优先查询，即只按这个条件查询。<br/>    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)<br/>    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。<br/>    * 入参fields传入images将查询多图数据，不传只返回主图数据。<br/>    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）<br/>    * 查询结果按照产品发布时间倒序，即时间近的数据在前。
*/
type TaobaoFenxiaoProductsGetRequest struct {
    model.Params
    // 商家编码
    outerId   string
    // 产品线ID
    productcatId   int64
    // 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    pids   string
    // 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    fields   string
    // 开始修改时间
    startModified   string
    // 结束修改时间
    endModified   string
    // 页码（大于0的整数，默认1）
    pageNo   int64
    // 每页记录数（默认20，最大50）
    pageSize   int64
    // sku商家编码
    skuNumber   string
    // 查询产品列表时，查询入参“是否需要授权”<br/>yes:需要授权 <br/>no:不需要授权
    isAuthz   string
    // 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    itemIds   string
}

// 初始化TaobaoFenxiaoProductsGetRequest对象
func NewTaobaoFenxiaoProductsGetRequest() *TaobaoFenxiaoProductsGetRequest{
    return &TaobaoFenxiaoProductsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductsGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.products.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 商家编码
func (r *TaobaoFenxiaoProductsGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoFenxiaoProductsGetRequest) GetOuterId() string {
    return r.outerId
}
// ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductsGetRequest) SetProductcatId(productcatId int64) error {
    r.productcatId = productcatId
    r.Set("productcat_id", productcatId)
    return nil
}

// ProductcatId Getter
func (r TaobaoFenxiaoProductsGetRequest) GetProductcatId() int64 {
    return r.productcatId
}
// Pids Setter
// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoProductsGetRequest) SetPids(pids string) error {
    r.pids = pids
    r.Set("pids", pids)
    return nil
}

// Pids Getter
func (r TaobaoFenxiaoProductsGetRequest) GetPids() string {
    return r.pids
}
// Fields Setter
// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (r *TaobaoFenxiaoProductsGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoProductsGetRequest) GetFields() string {
    return r.fields
}
// StartModified Setter
// 开始修改时间
func (r *TaobaoFenxiaoProductsGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

// StartModified Getter
func (r TaobaoFenxiaoProductsGetRequest) GetStartModified() string {
    return r.startModified
}
// EndModified Setter
// 结束修改时间
func (r *TaobaoFenxiaoProductsGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

// EndModified Getter
func (r TaobaoFenxiaoProductsGetRequest) GetEndModified() string {
    return r.endModified
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoProductsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoProductsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoProductsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoProductsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// SkuNumber Setter
// sku商家编码
func (r *TaobaoFenxiaoProductsGetRequest) SetSkuNumber(skuNumber string) error {
    r.skuNumber = skuNumber
    r.Set("sku_number", skuNumber)
    return nil
}

// SkuNumber Getter
func (r TaobaoFenxiaoProductsGetRequest) GetSkuNumber() string {
    return r.skuNumber
}
// IsAuthz Setter
// 查询产品列表时，查询入参“是否需要授权”<br/>yes:需要授权 <br/>no:不需要授权
func (r *TaobaoFenxiaoProductsGetRequest) SetIsAuthz(isAuthz string) error {
    r.isAuthz = isAuthz
    r.Set("is_authz", isAuthz)
    return nil
}

// IsAuthz Getter
func (r TaobaoFenxiaoProductsGetRequest) GetIsAuthz() string {
    return r.isAuthz
}
// ItemIds Setter
// 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoProductsGetRequest) SetItemIds(itemIds string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoFenxiaoProductsGetRequest) GetItemIds() string {
    return r.itemIds
}
