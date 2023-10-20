package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenPointOperateAPIRequest 积分操作接口 API请求
// alibaba.alsc.crm.open.point.operate
//
// 同步积分接口
type AlibabaAlscCrmOpenPointOperateAPIRequest struct {
	model.Params
	// 入参
	_paramPointOperateOpenReq *PointOperateOpenReq
}

// NewAlibabaAlscCrmOpenPointOperateRequest 初始化AlibabaAlscCrmOpenPointOperateAPIRequest对象
func NewAlibabaAlscCrmOpenPointOperateRequest() *AlibabaAlscCrmOpenPointOperateAPIRequest {
	return &AlibabaAlscCrmOpenPointOperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenPointOperateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.point.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenPointOperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenPointOperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPointOperateOpenReq is ParamPointOperateOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenPointOperateAPIRequest) SetParamPointOperateOpenReq(_paramPointOperateOpenReq *PointOperateOpenReq) error {
	r._paramPointOperateOpenReq = _paramPointOperateOpenReq
	r.Set("param_point_operate_open_req", _paramPointOperateOpenReq)
	return nil
}

// GetParamPointOperateOpenReq ParamPointOperateOpenReq Getter
func (r AlibabaAlscCrmOpenPointOperateAPIRequest) GetParamPointOperateOpenReq() *PointOperateOpenReq {
	return r._paramPointOperateOpenReq
}
