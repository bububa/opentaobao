package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员抵扣规则 APIRequest
alibaba.nlife.b2c.member.discountrule.get

获取企业会员抵扣规则
*/
type AlibabaNlifeB2cMemberDiscountruleGetRequest struct {
    model.Params

    // 企业ID
    companyId   string 

    // 会员在ISV处的编号
    cardNo   string 

}

func NewAlibabaNlifeB2cMemberDiscountruleGetRequest() *AlibabaNlifeB2cMemberDiscountruleGetRequest{
    return &AlibabaNlifeB2cMemberDiscountruleGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.member.discountrule.get"
}

func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeB2cMemberDiscountruleGetRequest) SetCompanyId(companyId string) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetCompanyId() string {
    return r.companyId
}

func (r *AlibabaNlifeB2cMemberDiscountruleGetRequest) SetCardNo(cardNo string) error {
    r.cardNo = cardNo
    r.Set("card_no", cardNo)
    return nil
}

func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetCardNo() string {
    return r.cardNo
}

