package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机设备信息查询 API请求
alibaba.lst.vending.equipment.query

为零售通品牌商提供已租赁的设备信息查询。
*/
type AlibabaLstVendingEquipmentQueryRequest struct {
    model.Params
    // 设备查询条件
    _openEquipmentQuery   *OpenEquipmentQuery
}

// 初始化AlibabaLstVendingEquipmentQueryRequest对象
func NewAlibabaLstVendingEquipmentQueryRequest() *AlibabaLstVendingEquipmentQueryRequest{
    return &AlibabaLstVendingEquipmentQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingEquipmentQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.equipment.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingEquipmentQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenEquipmentQuery Setter
// 设备查询条件
func (r *AlibabaLstVendingEquipmentQueryRequest) SetOpenEquipmentQuery(_openEquipmentQuery *OpenEquipmentQuery) error {
    r._openEquipmentQuery = _openEquipmentQuery
    r.Set("open_equipment_query", _openEquipmentQuery)
    return nil
}

// OpenEquipmentQuery Getter
func (r AlibabaLstVendingEquipmentQueryRequest) GetOpenEquipmentQuery() *OpenEquipmentQuery {
    return r._openEquipmentQuery
}
