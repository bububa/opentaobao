package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest 同步下架接口 API请求
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
type AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest struct {
	model.Params
	// 入参
	_paramSyncOffSaleCommandDTO *SyncOffSaleCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchSynoffsaleRequest 初始化AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchSynoffsaleRequest() *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.synoffsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncOffSaleCommandDTO is ParamSyncOffSaleCommandDTO Setter
// 入参
func (r *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest) SetParamSyncOffSaleCommandDTO(_paramSyncOffSaleCommandDTO *SyncOffSaleCommandDto) error {
	r._paramSyncOffSaleCommandDTO = _paramSyncOffSaleCommandDTO
	r.Set("param_sync_off_sale_command_d_t_o", _paramSyncOffSaleCommandDTO)
	return nil
}

// GetParamSyncOffSaleCommandDTO ParamSyncOffSaleCommandDTO Getter
func (r AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest) GetParamSyncOffSaleCommandDTO() *SyncOffSaleCommandDto {
	return r._paramSyncOffSaleCommandDTO
}
