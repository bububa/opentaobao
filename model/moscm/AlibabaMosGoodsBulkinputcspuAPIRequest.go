package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsBulkinputcspuAPIRequest 批量录入商品信息 API请求
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
type AlibabaMosGoodsBulkinputcspuAPIRequest struct {
	model.Params
	// 录入商品信息列表（最大列表长度：20）
	_cspuInputDtoList []CspuInputDto
}

// NewAlibabaMosGoodsBulkinputcspuRequest 初始化AlibabaMosGoodsBulkinputcspuAPIRequest对象
func NewAlibabaMosGoodsBulkinputcspuRequest() *AlibabaMosGoodsBulkinputcspuAPIRequest {
	return &AlibabaMosGoodsBulkinputcspuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.bulkinputcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCspuInputDtoList is CspuInputDtoList Setter
// 录入商品信息列表（最大列表长度：20）
func (r *AlibabaMosGoodsBulkinputcspuAPIRequest) SetCspuInputDtoList(_cspuInputDtoList []CspuInputDto) error {
	r._cspuInputDtoList = _cspuInputDtoList
	r.Set("cspu_input_dto_list", _cspuInputDtoList)
	return nil
}

// GetCspuInputDtoList CspuInputDtoList Getter
func (r AlibabaMosGoodsBulkinputcspuAPIRequest) GetCspuInputDtoList() []CspuInputDto {
	return r._cspuInputDtoList
}
