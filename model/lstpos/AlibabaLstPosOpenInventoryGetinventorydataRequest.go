package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品库存只读接口(最多20条库存信息) APIRequest
alibaba.lst.pos.open.inventory.getinventorydata

商品库存只读接口(最多20条库存信息)
*/
type AlibabaLstPosOpenInventoryGetinventorydataRequest struct {
    model.Params

    // ISV商品Id列表
    isvGoodsIdList   []string 

    // 门店主账号Id
    userId   int64 

}

func NewAlibabaLstPosOpenInventoryGetinventorydataRequest() *AlibabaLstPosOpenInventoryGetinventorydataRequest{
    return &AlibabaLstPosOpenInventoryGetinventorydataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.inventory.getinventorydata"
}

func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenInventoryGetinventorydataRequest) SetIsvGoodsIdList(isvGoodsIdList []string) error {
    r.isvGoodsIdList = isvGoodsIdList
    r.Set("isv_goods_id_list", isvGoodsIdList)
    return nil
}

func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetIsvGoodsIdList() []string {
    return r.isvGoodsIdList
}

func (r *AlibabaLstPosOpenInventoryGetinventorydataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenInventoryGetinventorydataRequest) GetUserId() int64 {
    return r.userId
}

