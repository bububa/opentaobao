package oversea

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOverseaExchagerateGetAPIRequest
汇率信息获取 API请求
alibaba.oversea.exchagerate.get

提供外部汇率查询接口 */
type AlibabaOverseaExchagerateGetAPIRequest struct {
	model.Params
	// 业务类型
	_bizCode string
	// 原始币种
	_baseCode string
	// 目标币种
	_targetCode string
}

// NewAlibabaOverseaExchagerateGetRequest 初始化AlibabaOverseaExchagerateGetAPIRequest对象
func NewAlibabaOverseaExchagerateGetRequest() *AlibabaOverseaExchagerateGetAPIRequest {
	return &AlibabaOverseaExchagerateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOverseaExchagerateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.exchagerate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOverseaExchagerateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizCode Setter
// 业务类型
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// Get BizCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetBizCode() string {
	return r._bizCode
}

// Set is BaseCode Setter
// 原始币种
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetBaseCode(_baseCode string) error {
	r._baseCode = _baseCode
	r.Set("base_code", _baseCode)
	return nil
}

// Get BaseCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetBaseCode() string {
	return r._baseCode
}

// Set is TargetCode Setter
// 目标币种
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetTargetCode(_targetCode string) error {
	r._targetCode = _targetCode
	r.Set("target_code", _targetCode)
	return nil
}

// Get TargetCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetTargetCode() string {
	return r._targetCode
}
