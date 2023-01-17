package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest 物流多包裹通知 API请求
// taobao.rdc.aligenius.logistics.packages.notice
//
// 订单发货之后，如果订单拆包、补发、赠品等场景，需要将多余包裹信息触达消费者, 大促会降级
type TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest struct {
	model.Params
	// 请求入参
	_paramLogisticsNoticeDTO *LogisticsNoticeDto
}

// NewTaobaoRdcAligeniusLogisticsPackagesNoticeRequest 初始化TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest对象
func NewTaobaoRdcAligeniusLogisticsPackagesNoticeRequest() *TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest {
	return &TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.logistics.packages.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLogisticsNoticeDTO is ParamLogisticsNoticeDTO Setter
// 请求入参
func (r *TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest) SetParamLogisticsNoticeDTO(_paramLogisticsNoticeDTO *LogisticsNoticeDto) error {
	r._paramLogisticsNoticeDTO = _paramLogisticsNoticeDTO
	r.Set("param_logistics_notice_d_t_o", _paramLogisticsNoticeDTO)
	return nil
}

// GetParamLogisticsNoticeDTO ParamLogisticsNoticeDTO Getter
func (r TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest) GetParamLogisticsNoticeDTO() *LogisticsNoticeDto {
	return r._paramLogisticsNoticeDTO
}
