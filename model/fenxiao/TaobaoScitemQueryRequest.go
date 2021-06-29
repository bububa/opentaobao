package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品 API请求
taobao.scitem.query

查询后端商品
*/
type TaobaoScitemQueryRequest struct {
    model.Params
    // 商品名称
    _itemName   string
    // 商家给商品的一个编码
    _outerCode   string
    // 条形码
    _barCode   string
    // ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
    _itemType   int64
    // 仓库编码
    _wmsCode   string
    // 当前页码数
    _pageIndex   int64
    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    _pageSize   int64
}

// 初始化TaobaoScitemQueryRequest对象
func NewTaobaoScitemQueryRequest() *TaobaoScitemQueryRequest{
    return &TaobaoScitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemQueryRequest) GetApiMethodName() string {
    return "taobao.scitem.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemName Setter
// 商品名称
func (r *TaobaoScitemQueryRequest) SetItemName(_itemName string) error {
    r._itemName = _itemName
    r.Set("item_name", _itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemQueryRequest) GetItemName() string {
    return r._itemName
}
// OuterCode Setter
// 商家给商品的一个编码
func (r *TaobaoScitemQueryRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemQueryRequest) GetOuterCode() string {
    return r._outerCode
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemQueryRequest) SetBarCode(_barCode string) error {
    r._barCode = _barCode
    r.Set("bar_code", _barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemQueryRequest) GetBarCode() string {
    return r._barCode
}
// ItemType Setter
// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
func (r *TaobaoScitemQueryRequest) SetItemType(_itemType int64) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemQueryRequest) GetItemType() int64 {
    return r._itemType
}
// WmsCode Setter
// 仓库编码
func (r *TaobaoScitemQueryRequest) SetWmsCode(_wmsCode string) error {
    r._wmsCode = _wmsCode
    r.Set("wms_code", _wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemQueryRequest) GetWmsCode() string {
    return r._wmsCode
}
// PageIndex Setter
// 当前页码数
func (r *TaobaoScitemQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoScitemQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoScitemQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoScitemQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
