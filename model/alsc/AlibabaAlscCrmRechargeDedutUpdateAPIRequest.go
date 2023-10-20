package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargededutupdateAPIRequest 储值消费 API请求
// alibaba.alsc.crm.recharge.dedut.update
//
// 增加储值消费接口
type AlibabaalsccrmrechargededutupdateAPIRequest struct {
	model.Params
	// 入参
	_paramDedutOpenReq *DedutOpenReq
}

// NewAlibabaalsccrmrechargededutupdateRequest 初始化AlibabaalsccrmrechargededutupdateAPIRequest对象
func NewAlibabaalsccrmrechargededutupdateRequest() *AlibabaalsccrmrechargededutupdateAPIRequest {
	return &AlibabaalsccrmrechargededutupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargededutupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.dedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargededutupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargededutupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDedutOpenReq is ParamDedutOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargededutupdateAPIRequest) SetParamDedutOpenReq(_paramDedutOpenReq *DedutOpenReq) error {
	r._paramDedutOpenReq = _paramDedutOpenReq
	r.Set("param_dedut_open_req", _paramDedutOpenReq)
	return nil
}

// GetParamDedutOpenReq ParamDedutOpenReq Getter
func (r AlibabaalsccrmrechargededutupdateAPIRequest) GetParamDedutOpenReq() *DedutOpenReq {
	return r._paramDedutOpenReq
}
