package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomequotationsyncAPIRequest 二手房行情数据同步 API请求
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
type AlibabaalihouseexistinghomequotationsyncAPIRequest struct {
	model.Params
	// dto
	_existingHouseQuotationReqDto *ExistingHouseQuotationReqDto
}

// NewAlibabaalihouseexistinghomequotationsyncRequest 初始化AlibabaalihouseexistinghomequotationsyncAPIRequest对象
func NewAlibabaalihouseexistinghomequotationsyncRequest() *AlibabaalihouseexistinghomequotationsyncAPIRequest {
	return &AlibabaalihouseexistinghomequotationsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomequotationsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.quotation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomequotationsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomequotationsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExistingHouseQuotationReqDto is ExistingHouseQuotationReqDto Setter
// dto
func (r *AlibabaalihouseexistinghomequotationsyncAPIRequest) SetExistingHouseQuotationReqDto(_existingHouseQuotationReqDto *ExistingHouseQuotationReqDto) error {
	r._existingHouseQuotationReqDto = _existingHouseQuotationReqDto
	r.Set("existing_house_quotation_req_dto", _existingHouseQuotationReqDto)
	return nil
}

// GetExistingHouseQuotationReqDto ExistingHouseQuotationReqDto Getter
func (r AlibabaalihouseexistinghomequotationsyncAPIRequest) GetExistingHouseQuotationReqDto() *ExistingHouseQuotationReqDto {
	return r._existingHouseQuotationReqDto
}
