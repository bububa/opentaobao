package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存修改同步接口(最多20条库存信息) API请求
alibaba.lst.pos.open.inventory.syncinventorydata

商品库存修改同步接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventorySyncinventorydataAPIRequest struct {
    model.Params
    // 库存对象列表
    _inventoryDTOList   []InventoryDTO
    // 门店对应的主账号id
    _userId   int64
}

// 初始化AlibabaLstPosOpenInventorySyncinventorydataAPIRequest对象
func NewAlibabaLstPosOpenInventorySyncinventorydataRequest() *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest{
    return &AlibabaLstPosOpenInventorySyncinventorydataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.inventory.syncinventorydata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InventoryDTOList Setter
// 库存对象列表
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetInventoryDTOList(_inventoryDTOList []InventoryDTO) error {
    r._inventoryDTOList = _inventoryDTOList
    r.Set("inventory_d_t_o_list", _inventoryDTOList)
    return nil
}

// InventoryDTOList Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetInventoryDTOList() []InventoryDTO {
    return r._inventoryDTOList
}
// UserId Setter
// 门店对应的主账号id
func (r *AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenInventorySyncinventorydataAPIRequest) GetUserId() int64 {
    return r._userId
}
