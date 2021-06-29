package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证批量申请人送签接口 API请求
alitrip.travel.visa.sign.send

签证批量申请人送签接口，用于商家批量送签。
*/
type AlitripTravelVisaSignSendRequest struct {
    model.Params
    // 国家id。目前只支持越南，越南国家id:27027
    _nationId   int64
    // 送签类型：1-非加急，2-加急，默认非加急
    _signType   int64
    // 申请人ids
    _applyIds   []string
}

// 初始化AlitripTravelVisaSignSendRequest对象
func NewAlitripTravelVisaSignSendRequest() *AlitripTravelVisaSignSendRequest{
    return &AlitripTravelVisaSignSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaSignSendRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.sign.send"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaSignSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NationId Setter
// 国家id。目前只支持越南，越南国家id:27027
func (r *AlitripTravelVisaSignSendRequest) SetNationId(_nationId int64) error {
    r._nationId = _nationId
    r.Set("nation_id", _nationId)
    return nil
}

// NationId Getter
func (r AlitripTravelVisaSignSendRequest) GetNationId() int64 {
    return r._nationId
}
// SignType Setter
// 送签类型：1-非加急，2-加急，默认非加急
func (r *AlitripTravelVisaSignSendRequest) SetSignType(_signType int64) error {
    r._signType = _signType
    r.Set("sign_type", _signType)
    return nil
}

// SignType Getter
func (r AlitripTravelVisaSignSendRequest) GetSignType() int64 {
    return r._signType
}
// ApplyIds Setter
// 申请人ids
func (r *AlitripTravelVisaSignSendRequest) SetApplyIds(_applyIds []string) error {
    r._applyIds = _applyIds
    r.Set("apply_ids", _applyIds)
    return nil
}

// ApplyIds Getter
func (r AlitripTravelVisaSignSendRequest) GetApplyIds() []string {
    return r._applyIds
}
