package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员抵扣规则 API请求
alibaba.nlife.b2c.member.discountrule.get

获取企业会员抵扣规则
*/
type AlibabaNlifeB2cMemberDiscountruleGetRequest struct {
    model.Params
    // 企业ID
    _companyId   string
    // 会员在ISV处的编号
    _cardNo   string
}

// 初始化AlibabaNlifeB2cMemberDiscountruleGetRequest对象
func NewAlibabaNlifeB2cMemberDiscountruleGetRequest() *AlibabaNlifeB2cMemberDiscountruleGetRequest{
    return &AlibabaNlifeB2cMemberDiscountruleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.member.discountrule.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 企业ID
func (r *AlibabaNlifeB2cMemberDiscountruleGetRequest) SetCompanyId(_companyId string) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetCompanyId() string {
    return r._companyId
}
// CardNo Setter
// 会员在ISV处的编号
func (r *AlibabaNlifeB2cMemberDiscountruleGetRequest) SetCardNo(_cardNo string) error {
    r._cardNo = _cardNo
    r.Set("card_no", _cardNo)
    return nil
}

// CardNo Getter
func (r AlibabaNlifeB2cMemberDiscountruleGetRequest) GetCardNo() string {
    return r._cardNo
}
