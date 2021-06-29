package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证批量申请人送签接口 APIRequest
alitrip.travel.visa.sign.send

签证批量申请人送签接口，用于商家批量送签。
*/
type AlitripTravelVisaSignSendRequest struct {
    model.Params

    // 国家id。目前只支持越南，越南国家id:27027
    nationId   int64 

    // 送签类型：1-非加急，2-加急，默认非加急
    signType   int64 

    // 申请人ids
    applyIds   []string 

}

func NewAlitripTravelVisaSignSendRequest() *AlitripTravelVisaSignSendRequest{
    return &AlitripTravelVisaSignSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelVisaSignSendRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.sign.send"
}

func (r AlitripTravelVisaSignSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelVisaSignSendRequest) SetNationId(nationId int64) error {
    r.nationId = nationId
    r.Set("nation_id", nationId)
    return nil
}

func (r AlitripTravelVisaSignSendRequest) GetNationId() int64 {
    return r.nationId
}

func (r *AlitripTravelVisaSignSendRequest) SetSignType(signType int64) error {
    r.signType = signType
    r.Set("sign_type", signType)
    return nil
}

func (r AlitripTravelVisaSignSendRequest) GetSignType() int64 {
    return r.signType
}

func (r *AlitripTravelVisaSignSendRequest) SetApplyIds(applyIds []string) error {
    r.applyIds = applyIds
    r.Set("apply_ids", applyIds)
    return nil
}

func (r AlitripTravelVisaSignSendRequest) GetApplyIds() []string {
    return r.applyIds
}

