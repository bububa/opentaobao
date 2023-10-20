package nlife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cMemberDiscountruleGetAPIRequest 会员抵扣规则 API请求
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
type AlibabaNlifeB2cMemberDiscountruleGetAPIRequest struct {
	model.Params
	// 企业ID
	_companyId string
	// 会员在ISV处的编号
	_cardNo string
}

// NewAlibabaNlifeB2cMemberDiscountruleGetRequest 初始化AlibabaNlifeB2cMemberDiscountruleGetAPIRequest对象
func NewAlibabaNlifeB2cMemberDiscountruleGetRequest() *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest {
	return &AlibabaNlifeB2cMemberDiscountruleGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) Reset() {
	r._companyId = ""
	r._cardNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.member.discountrule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyId is CompanyId Setter
// 企业ID
func (r *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) SetCompanyId(_companyId string) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) GetCompanyId() string {
	return r._companyId
}

// SetCardNo is CardNo Setter
// 会员在ISV处的编号
func (r *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) GetCardNo() string {
	return r._cardNo
}

var poolAlibabaNlifeB2cMemberDiscountruleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cMemberDiscountruleGetRequest()
	},
}

// GetAlibabaNlifeB2cMemberDiscountruleGetRequest 从 sync.Pool 获取 AlibabaNlifeB2cMemberDiscountruleGetAPIRequest
func GetAlibabaNlifeB2cMemberDiscountruleGetAPIRequest() *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest {
	return poolAlibabaNlifeB2cMemberDiscountruleGetAPIRequest.Get().(*AlibabaNlifeB2cMemberDiscountruleGetAPIRequest)
}

// ReleaseAlibabaNlifeB2cMemberDiscountruleGetAPIRequest 将 AlibabaNlifeB2cMemberDiscountruleGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cMemberDiscountruleGetAPIRequest(v *AlibabaNlifeB2cMemberDiscountruleGetAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cMemberDiscountruleGetAPIRequest.Put(v)
}
