package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息查询接口 APIRequest
alibaba.westcrm.customer.info.get

会员信息查询接口
*/
type AlibabaWestcrmCustomerInfoGetRequest struct {
    model.Params

    // 园区id
    campusId   int64 

    // 会员id
    ibUserId   int64 

    // 支付宝id
    alipayId   string 

}

func NewAlibabaWestcrmCustomerInfoGetRequest() *AlibabaWestcrmCustomerInfoGetRequest{
    return &AlibabaWestcrmCustomerInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmCustomerInfoGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.customer.info.get"
}

func (r AlibabaWestcrmCustomerInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmCustomerInfoGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmCustomerInfoGetRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmCustomerInfoGetRequest) SetIbUserId(ibUserId int64) error {
    r.ibUserId = ibUserId
    r.Set("ib_user_id", ibUserId)
    return nil
}

func (r AlibabaWestcrmCustomerInfoGetRequest) GetIbUserId() int64 {
    return r.ibUserId
}

func (r *AlibabaWestcrmCustomerInfoGetRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaWestcrmCustomerInfoGetRequest) GetAlipayId() string {
    return r.alipayId
}

