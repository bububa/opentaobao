package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest 同步下架接口 API请求
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
type AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest struct {
	model.Params
	// 入参
	_paramSyncOffSaleCommandDTO *SyncOffSaleCommandDto
}

// NewAlibabajymitemexternalgoodsbatchsynoffsaleRequest 初始化AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchsynoffsaleRequest() *AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest {
	return &AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.synoffsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncOffSaleCommandDTO is ParamSyncOffSaleCommandDTO Setter
// 入参
func (r *AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest) SetParamSyncOffSaleCommandDTO(_paramSyncOffSaleCommandDTO *SyncOffSaleCommandDto) error {
	r._paramSyncOffSaleCommandDTO = _paramSyncOffSaleCommandDTO
	r.Set("param_sync_off_sale_command_d_t_o", _paramSyncOffSaleCommandDTO)
	return nil
}

// GetParamSyncOffSaleCommandDTO ParamSyncOffSaleCommandDTO Getter
func (r AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest) GetParamSyncOffSaleCommandDTO() *SyncOffSaleCommandDto {
	return r._paramSyncOffSaleCommandDTO
}
