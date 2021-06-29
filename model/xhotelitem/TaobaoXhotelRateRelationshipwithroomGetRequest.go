package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rpId APIRequest
taobao.xhotel.rate.relationshipwithroom.get

某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
*/
type TaobaoXhotelRateRelationshipwithroomGetRequest struct {
    model.Params

    // rpId
    rpId   int64 

    // 页数
    pageNo   int64 

}

func NewTaobaoXhotelRateRelationshipwithroomGetRequest() *TaobaoXhotelRateRelationshipwithroomGetRequest{
    return &TaobaoXhotelRateRelationshipwithroomGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.relationshipwithroom.get"
}

func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateRelationshipwithroomGetRequest) SetRpId(rpId int64) error {
    r.rpId = rpId
    r.Set("rp_id", rpId)
    return nil
}

func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetRpId() int64 {
    return r.rpId
}

func (r *TaobaoXhotelRateRelationshipwithroomGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetPageNo() int64 {
    return r.pageNo
}

