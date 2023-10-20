package oversea

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOverseaExchagerateGetAPIRequest 汇率信息获取 API请求
// alibaba.oversea.exchagerate.get
//
// 提供外部汇率查询接口
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOverseaExchagerateGetAPIRequest) Reset() {
	r._bizCode = ""
	r._baseCode = ""
	r._targetCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOverseaExchagerateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.exchagerate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOverseaExchagerateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOverseaExchagerateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizCode is BizCode Setter
// 业务类型
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetBaseCode is BaseCode Setter
// 原始币种
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetBaseCode(_baseCode string) error {
	r._baseCode = _baseCode
	r.Set("base_code", _baseCode)
	return nil
}

// GetBaseCode BaseCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetBaseCode() string {
	return r._baseCode
}

// SetTargetCode is TargetCode Setter
// 目标币种
func (r *AlibabaOverseaExchagerateGetAPIRequest) SetTargetCode(_targetCode string) error {
	r._targetCode = _targetCode
	r.Set("target_code", _targetCode)
	return nil
}

// GetTargetCode TargetCode Getter
func (r AlibabaOverseaExchagerateGetAPIRequest) GetTargetCode() string {
	return r._targetCode
}

var poolAlibabaOverseaExchagerateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOverseaExchagerateGetRequest()
	},
}

// GetAlibabaOverseaExchagerateGetRequest 从 sync.Pool 获取 AlibabaOverseaExchagerateGetAPIRequest
func GetAlibabaOverseaExchagerateGetAPIRequest() *AlibabaOverseaExchagerateGetAPIRequest {
	return poolAlibabaOverseaExchagerateGetAPIRequest.Get().(*AlibabaOverseaExchagerateGetAPIRequest)
}

// ReleaseAlibabaOverseaExchagerateGetAPIRequest 将 AlibabaOverseaExchagerateGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaOverseaExchagerateGetAPIRequest(v *AlibabaOverseaExchagerateGetAPIRequest) {
	v.Reset()
	poolAlibabaOverseaExchagerateGetAPIRequest.Put(v)
}
