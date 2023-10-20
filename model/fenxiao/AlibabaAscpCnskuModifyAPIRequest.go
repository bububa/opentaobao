package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpCnskuModifyAPIRequest 供应链中台货品修改接口 API请求
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
type AlibabaAscpCnskuModifyAPIRequest struct {
	model.Params
	// 待修改的货品
	_cnskuDto *CnskuDto
	// 选项
	_updateCnskuOption *UpdateCnskuOption
}

// NewAlibabaAscpCnskuModifyRequest 初始化AlibabaAscpCnskuModifyAPIRequest对象
func NewAlibabaAscpCnskuModifyRequest() *AlibabaAscpCnskuModifyAPIRequest {
	return &AlibabaAscpCnskuModifyAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpCnskuModifyAPIRequest) Reset() {
	r._cnskuDto = nil
	r._updateCnskuOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.cnsku.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpCnskuModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnskuDto is CnskuDto Setter
// 待修改的货品
func (r *AlibabaAscpCnskuModifyAPIRequest) SetCnskuDto(_cnskuDto *CnskuDto) error {
	r._cnskuDto = _cnskuDto
	r.Set("cnsku_dto", _cnskuDto)
	return nil
}

// GetCnskuDto CnskuDto Getter
func (r AlibabaAscpCnskuModifyAPIRequest) GetCnskuDto() *CnskuDto {
	return r._cnskuDto
}

// SetUpdateCnskuOption is UpdateCnskuOption Setter
// 选项
func (r *AlibabaAscpCnskuModifyAPIRequest) SetUpdateCnskuOption(_updateCnskuOption *UpdateCnskuOption) error {
	r._updateCnskuOption = _updateCnskuOption
	r.Set("update_cnsku_option", _updateCnskuOption)
	return nil
}

// GetUpdateCnskuOption UpdateCnskuOption Getter
func (r AlibabaAscpCnskuModifyAPIRequest) GetUpdateCnskuOption() *UpdateCnskuOption {
	return r._updateCnskuOption
}

var poolAlibabaAscpCnskuModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpCnskuModifyRequest()
	},
}

// GetAlibabaAscpCnskuModifyRequest 从 sync.Pool 获取 AlibabaAscpCnskuModifyAPIRequest
func GetAlibabaAscpCnskuModifyAPIRequest() *AlibabaAscpCnskuModifyAPIRequest {
	return poolAlibabaAscpCnskuModifyAPIRequest.Get().(*AlibabaAscpCnskuModifyAPIRequest)
}

// ReleaseAlibabaAscpCnskuModifyAPIRequest 将 AlibabaAscpCnskuModifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpCnskuModifyAPIRequest(v *AlibabaAscpCnskuModifyAPIRequest) {
	v.Reset()
	poolAlibabaAscpCnskuModifyAPIRequest.Put(v)
}
