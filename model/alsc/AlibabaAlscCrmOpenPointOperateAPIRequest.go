package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenpointoperateAPIRequest 积分操作接口 API请求
// alibaba.alsc.crm.open.point.operate
//
// 同步积分接口
type AlibabaalsccrmopenpointoperateAPIRequest struct {
	model.Params
	// 入参
	_paramPointOperateOpenReq *PointOperateOpenReq
}

// NewAlibabaalsccrmopenpointoperateRequest 初始化AlibabaalsccrmopenpointoperateAPIRequest对象
func NewAlibabaalsccrmopenpointoperateRequest() *AlibabaalsccrmopenpointoperateAPIRequest {
	return &AlibabaalsccrmopenpointoperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenpointoperateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.point.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenpointoperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenpointoperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPointOperateOpenReq is ParamPointOperateOpenReq Setter
// 入参
func (r *AlibabaalsccrmopenpointoperateAPIRequest) SetParamPointOperateOpenReq(_paramPointOperateOpenReq *PointOperateOpenReq) error {
	r._paramPointOperateOpenReq = _paramPointOperateOpenReq
	r.Set("param_point_operate_open_req", _paramPointOperateOpenReq)
	return nil
}

// GetParamPointOperateOpenReq ParamPointOperateOpenReq Getter
func (r AlibabaalsccrmopenpointoperateAPIRequest) GetParamPointOperateOpenReq() *PointOperateOpenReq {
	return r._paramPointOperateOpenReq
}
