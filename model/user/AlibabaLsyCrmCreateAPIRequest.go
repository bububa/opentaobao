package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmCreateAPIRequest
创建客资 API请求
alibaba.lsy.crm.create

欢客调用该接口进行客资创建 */
type AlibabaLsyCrmCreateAPIRequest struct {
	model.Params
	// 客资记录对象
	_nrtRecordDto *NrtRecordDto
}

// NewAlibabaLsyCrmCreateRequest 初始化AlibabaLsyCrmCreateAPIRequest对象
func NewAlibabaLsyCrmCreateRequest() *AlibabaLsyCrmCreateAPIRequest {
	return &AlibabaLsyCrmCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NrtRecordDto Setter
// 客资记录对象
func (r *AlibabaLsyCrmCreateAPIRequest) SetNrtRecordDto(_nrtRecordDto *NrtRecordDto) error {
	r._nrtRecordDto = _nrtRecordDto
	r.Set("nrt_record_dto", _nrtRecordDto)
	return nil
}

// Get NrtRecordDto Getter
func (r AlibabaLsyCrmCreateAPIRequest) GetNrtRecordDto() *NrtRecordDto {
	return r._nrtRecordDto
}
