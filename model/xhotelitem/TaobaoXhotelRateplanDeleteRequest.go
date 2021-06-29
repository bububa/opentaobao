package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan删除 APIRequest
taobao.xhotel.rateplan.delete

酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
*/
type TaobaoXhotelRateplanDeleteRequest struct {
    model.Params

    // ratepland标识
    rpId   int64 

    // 系统商，一般不用填写，使用须申请
    vendor   string 

    // 商家价格政策编码
    rateplanCode   string 

}

func NewTaobaoXhotelRateplanDeleteRequest() *TaobaoXhotelRateplanDeleteRequest{
    return &TaobaoXhotelRateplanDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateplanDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.delete"
}

func (r TaobaoXhotelRateplanDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateplanDeleteRequest) SetRpId(rpId int64) error {
    r.rpId = rpId
    r.Set("rp_id", rpId)
    return nil
}

func (r TaobaoXhotelRateplanDeleteRequest) GetRpId() int64 {
    return r.rpId
}

func (r *TaobaoXhotelRateplanDeleteRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateplanDeleteRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRateplanDeleteRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateplanDeleteRequest) GetRateplanCode() string {
    return r.rateplanCode
}

