package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存修改同步接口(最多20条库存信息) APIRequest
alibaba.lst.pos.open.inventory.syncinventorydata

商品库存修改同步接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventorySyncinventorydataRequest struct {
    model.Params

    // 库存对象列表
    inventoryDTOList   []InventoryDto 

    // 门店对应的主账号id
    userId   int64 

}

func NewAlibabaLstPosOpenInventorySyncinventorydataRequest() *AlibabaLstPosOpenInventorySyncinventorydataRequest{
    return &AlibabaLstPosOpenInventorySyncinventorydataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenInventorySyncinventorydataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.inventory.syncinventorydata"
}

func (r AlibabaLstPosOpenInventorySyncinventorydataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenInventorySyncinventorydataRequest) SetInventoryDTOList(inventoryDTOList []InventoryDto) error {
    r.inventoryDTOList = inventoryDTOList
    r.Set("inventory_d_t_o_list", inventoryDTOList)
    return nil
}

func (r AlibabaLstPosOpenInventorySyncinventorydataRequest) GetInventoryDTOList() []InventoryDto {
    return r.inventoryDTOList
}

func (r *AlibabaLstPosOpenInventorySyncinventorydataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenInventorySyncinventorydataRequest) GetUserId() int64 {
    return r.userId
}

