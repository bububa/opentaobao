package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取交易列表 APIRequest
taobao.tanx.deals.get

批量获取交易信息
*/
type TaobaoTanxDealsGetRequest struct {
    model.Params

    // dsp用户id
    dspId   int64 

    // dsp用户验证token
    token   string 

    // 页大小
    pageSize   int64 

    // 交易类型
    dealType   int64 

    // 页码
    page   int64 

    // 1970年到现在的时间，毫秒
    signTime   int64 

}

func NewTaobaoTanxDealsGetRequest() *TaobaoTanxDealsGetRequest{
    return &TaobaoTanxDealsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxDealsGetRequest) GetApiMethodName() string {
    return "taobao.tanx.deals.get"
}

func (r TaobaoTanxDealsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxDealsGetRequest) SetDspId(dspId int64) error {
    r.dspId = dspId
    r.Set("dsp_id", dspId)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetDspId() int64 {
    return r.dspId
}

func (r *TaobaoTanxDealsGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxDealsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTanxDealsGetRequest) SetDealType(dealType int64) error {
    r.dealType = dealType
    r.Set("deal_type", dealType)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetDealType() int64 {
    return r.dealType
}

func (r *TaobaoTanxDealsGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetPage() int64 {
    return r.page
}

func (r *TaobaoTanxDealsGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxDealsGetRequest) GetSignTime() int64 {
    return r.signTime
}

