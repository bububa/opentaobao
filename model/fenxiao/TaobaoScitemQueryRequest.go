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
    itemName   string
    // 商家给商品的一个编码
    outerCode   string
    // 条形码
    barCode   string
    // ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
    itemType   int64
    // 仓库编码
    wmsCode   string
    // 当前页码数
    pageIndex   int64
    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    pageSize   int64
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
func (r *TaobaoScitemQueryRequest) SetItemName(itemName string) error {
    r.itemName = itemName
    r.Set("item_name", itemName)
    return nil
}

// ItemName Getter
func (r TaobaoScitemQueryRequest) GetItemName() string {
    return r.itemName
}
// OuterCode Setter
// 商家给商品的一个编码
func (r *TaobaoScitemQueryRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemQueryRequest) GetOuterCode() string {
    return r.outerCode
}
// BarCode Setter
// 条形码
func (r *TaobaoScitemQueryRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

// BarCode Getter
func (r TaobaoScitemQueryRequest) GetBarCode() string {
    return r.barCode
}
// ItemType Setter
// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
func (r *TaobaoScitemQueryRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r TaobaoScitemQueryRequest) GetItemType() int64 {
    return r.itemType
}
// WmsCode Setter
// 仓库编码
func (r *TaobaoScitemQueryRequest) SetWmsCode(wmsCode string) error {
    r.wmsCode = wmsCode
    r.Set("wms_code", wmsCode)
    return nil
}

// WmsCode Getter
func (r TaobaoScitemQueryRequest) GetWmsCode() string {
    return r.wmsCode
}
// PageIndex Setter
// 当前页码数
func (r *TaobaoScitemQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoScitemQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoScitemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoScitemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
