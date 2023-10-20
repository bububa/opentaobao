package oversea

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaoverseaexchagerategetAPIRequest 汇率信息获取 API请求
// alibaba.oversea.exchagerate.get
//
// 提供外部汇率查询接口
type AlibabaoverseaexchagerategetAPIRequest struct {
	model.Params
	// 业务类型
	_bizCode string
	// 原始币种
	_baseCode string
	// 目标币种
	_targetCode string
}

// NewAlibabaoverseaexchagerategetRequest 初始化AlibabaoverseaexchagerategetAPIRequest对象
func NewAlibabaoverseaexchagerategetRequest() *AlibabaoverseaexchagerategetAPIRequest {
	return &AlibabaoverseaexchagerategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaoverseaexchagerategetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.exchagerate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaoverseaexchagerategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaoverseaexchagerategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务类型
func (r *AlibabaoverseaexchagerategetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AlibabaoverseaexchagerategetAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetBaseCode is BaseCode Setter
// 原始币种
func (r *AlibabaoverseaexchagerategetAPIRequest) SetBaseCode(_baseCode string) error {
	r._baseCode = _baseCode
	r.Set("base_code", _baseCode)
	return nil
}

// GetBaseCode BaseCode Getter
func (r AlibabaoverseaexchagerategetAPIRequest) GetBaseCode() string {
	return r._baseCode
}

// SetTargetCode is TargetCode Setter
// 目标币种
func (r *AlibabaoverseaexchagerategetAPIRequest) SetTargetCode(_targetCode string) error {
	r._targetCode = _targetCode
	r.Set("target_code", _targetCode)
	return nil
}

// GetTargetCode TargetCode Getter
func (r AlibabaoverseaexchagerategetAPIRequest) GetTargetCode() string {
	return r._targetCode
}
