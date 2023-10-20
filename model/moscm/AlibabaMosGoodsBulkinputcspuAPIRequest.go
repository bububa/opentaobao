package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodsbulkinputcspuAPIRequest 批量录入商品信息 API请求
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
type AlibabamosgoodsbulkinputcspuAPIRequest struct {
	model.Params
	// 录入商品信息列表（最大列表长度：20）
	_cspuInputDtoList []CspuInputDto
}

// NewAlibabamosgoodsbulkinputcspuRequest 初始化AlibabamosgoodsbulkinputcspuAPIRequest对象
func NewAlibabamosgoodsbulkinputcspuRequest() *AlibabamosgoodsbulkinputcspuAPIRequest {
	return &AlibabamosgoodsbulkinputcspuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodsbulkinputcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.bulkinputcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodsbulkinputcspuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodsbulkinputcspuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCspuInputDtoList is CspuInputDtoList Setter
// 录入商品信息列表（最大列表长度：20）
func (r *AlibabamosgoodsbulkinputcspuAPIRequest) SetCspuInputDtoList(_cspuInputDtoList []CspuInputDto) error {
	r._cspuInputDtoList = _cspuInputDtoList
	r.Set("cspu_input_dto_list", _cspuInputDtoList)
	return nil
}

// GetCspuInputDtoList CspuInputDtoList Getter
func (r AlibabamosgoodsbulkinputcspuAPIRequest) GetCspuInputDtoList() []CspuInputDto {
	return r._cspuInputDtoList
}
