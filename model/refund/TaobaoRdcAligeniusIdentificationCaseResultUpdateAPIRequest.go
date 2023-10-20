package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniusidentificationcaseresultupdateAPIRequest 鉴定工单结果同步 API请求
// taobao.rdc.aligenius.identification.case.result.update
//
// 同步鉴定工单结果信息
type TaobaordcaligeniusidentificationcaseresultupdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *SyncIdentifyRefundCaseResultDto
}

// NewTaobaordcaligeniusidentificationcaseresultupdateRequest 初始化TaobaordcaligeniusidentificationcaseresultupdateAPIRequest对象
func NewTaobaordcaligeniusidentificationcaseresultupdateRequest() *TaobaordcaligeniusidentificationcaseresultupdateAPIRequest {
	return &TaobaordcaligeniusidentificationcaseresultupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniusidentificationcaseresultupdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.identification.case.result.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniusidentificationcaseresultupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniusidentificationcaseresultupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *TaobaordcaligeniusidentificationcaseresultupdateAPIRequest) SetParam(_param *SyncIdentifyRefundCaseResultDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaordcaligeniusidentificationcaseresultupdateAPIRequest) GetParam() *SyncIdentifyRefundCaseResultDto {
	return r._param
}
