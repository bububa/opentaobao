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
type TaobaoFenxiaoProductsGetAPIRequest struct {
    model.Params
    // 商家编码
    _outerId   string
    // 产品线ID
    _productcatId   int64
    // 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
    _pids   string
    // 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
    _fields   string
    // 开始修改时间
    _startModified   string
    // 结束修改时间
    _endModified   string
    // 页码（大于0的整数，默认1）
    _pageNo   int64
    // 每页记录数（默认20，最大50）
    _pageSize   int64
    // sku商家编码
    _skuNumber   string
    // 查询产品列表时，查询入参“是否需要授权”<br/>yes:需要授权 <br/>no:不需要授权
    _isAuthz   string
    // 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
    _itemIds   string
}

// 初始化TaobaoFenxiaoProductsGetAPIRequest对象
func NewTaobaoFenxiaoProductsGetRequest() *TaobaoFenxiaoProductsGetAPIRequest{
    return &TaobaoFenxiaoProductsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductsGetAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.products.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 商家编码
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetOuterId() string {
    return r._outerId
}
// ProductcatId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetProductcatId(_productcatId int64) error {
    r._productcatId = _productcatId
    r.Set("productcat_id", _productcatId)
    return nil
}

// ProductcatId Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetProductcatId() int64 {
    return r._productcatId
}
// Pids Setter
// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetPids(_pids string) error {
    r._pids = _pids
    r.Set("pids", _pids)
    return nil
}

// Pids Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetPids() string {
    return r._pids
}
// Fields Setter
// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetFields() string {
    return r._fields
}
// StartModified Setter
// 开始修改时间
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 结束修改时间
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetEndModified() string {
    return r._endModified
}
// PageNo Setter
// 页码（大于0的整数，默认1）
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数（默认20，最大50）
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// SkuNumber Setter
// sku商家编码
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetSkuNumber(_skuNumber string) error {
    r._skuNumber = _skuNumber
    r.Set("sku_number", _skuNumber)
    return nil
}

// SkuNumber Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetSkuNumber() string {
    return r._skuNumber
}
// IsAuthz Setter
// 查询产品列表时，查询入参“是否需要授权”<br/>yes:需要授权 <br/>no:不需要授权
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetIsAuthz(_isAuthz string) error {
    r._isAuthz = _isAuthz
    r.Set("is_authz", _isAuthz)
    return nil
}

// IsAuthz Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetIsAuthz() string {
    return r._isAuthz
}
// ItemIds Setter
// 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (r *TaobaoFenxiaoProductsGetAPIRequest) SetItemIds(_itemIds string) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoFenxiaoProductsGetAPIRequest) GetItemIds() string {
    return r._itemIds
}
