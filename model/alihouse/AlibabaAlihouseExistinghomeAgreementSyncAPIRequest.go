package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeagreementsyncAPIRequest 二手房电子协议数据同步 API请求
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
type AlibabaalihouseexistinghomeagreementsyncAPIRequest struct {
	model.Params
	// 数据结构
	_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto
}

// NewAlibabaalihouseexistinghomeagreementsyncRequest 初始化AlibabaalihouseexistinghomeagreementsyncAPIRequest对象
func NewAlibabaalihouseexistinghomeagreementsyncRequest() *AlibabaalihouseexistinghomeagreementsyncAPIRequest {
	return &AlibabaalihouseexistinghomeagreementsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeagreementsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.agreement.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeagreementsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeagreementsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExistingHomeElectricAgreementDto is ExistingHomeElectricAgreementDto Setter
// 数据结构
func (r *AlibabaalihouseexistinghomeagreementsyncAPIRequest) SetExistingHomeElectricAgreementDto(_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto) error {
	r._existingHomeElectricAgreementDto = _existingHomeElectricAgreementDto
	r.Set("existing_home_electric_agreement_dto", _existingHomeElectricAgreementDto)
	return nil
}

// GetExistingHomeElectricAgreementDto ExistingHomeElectricAgreementDto Getter
func (r AlibabaalihouseexistinghomeagreementsyncAPIRequest) GetExistingHomeElectricAgreementDto() *ExistingHomeElectricAgreementDto {
	return r._existingHomeElectricAgreementDto
}
