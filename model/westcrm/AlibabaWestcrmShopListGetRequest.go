package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商圈商户信息列表 APIRequest
alibaba.westcrm.shop.list.get

获取商圈商户信息列表
*/
type AlibabaWestcrmShopListGetRequest struct {
    model.Params

    // 园区id
    campusId   int64 

}

func NewAlibabaWestcrmShopListGetRequest() *AlibabaWestcrmShopListGetRequest{
    return &AlibabaWestcrmShopListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmShopListGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.shop.list.get"
}

func (r AlibabaWestcrmShopListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmShopListGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmShopListGetRequest) GetCampusId() int64 {
    return r.campusId
}

