package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询rpId API请求
taobao.xhotel.rate.relationshipwithroom.get

某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
*/
type TaobaoXhotelRateRelationshipwithroomGetRequest struct {
    model.Params
    // rpId
    _rpId   int64
    // 页数
    _pageNo   int64
}

// 初始化TaobaoXhotelRateRelationshipwithroomGetRequest对象
func NewTaobaoXhotelRateRelationshipwithroomGetRequest() *TaobaoXhotelRateRelationshipwithroomGetRequest{
    return &TaobaoXhotelRateRelationshipwithroomGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.relationshipwithroom.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RpId Setter
// rpId
func (r *TaobaoXhotelRateRelationshipwithroomGetRequest) SetRpId(_rpId int64) error {
    r._rpId = _rpId
    r.Set("rp_id", _rpId)
    return nil
}

// RpId Getter
func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetRpId() int64 {
    return r._rpId
}
// PageNo Setter
// 页数
func (r *TaobaoXhotelRateRelationshipwithroomGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoXhotelRateRelationshipwithroomGetRequest) GetPageNo() int64 {
    return r._pageNo
}
