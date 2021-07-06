package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperCommonReturnAPIRequest 纸票通用回传接口 API请求
// alibaba.einvoice.paper.common.return
//
// 纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
type AlibabaEinvoicePaperCommonReturnAPIRequest struct {
	model.Params
	// 请求索引
	_reqIndex string
	// 错误码，success=false时必填
	_bizErrorCode string
	// 错误信息，success=false时必填
	_bizErrorMsg string
	// 扩展信息
	_extProps string
	// 回传结果
	_success bool
}

// NewAlibabaEinvoicePaperCommonReturnRequest 初始化AlibabaEinvoicePaperCommonReturnAPIRequest对象
func NewAlibabaEinvoicePaperCommonReturnRequest() *AlibabaEinvoicePaperCommonReturnAPIRequest {
	return &AlibabaEinvoicePaperCommonReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.common.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReqIndex is ReqIndex Setter
// 请求索引
func (r *AlibabaEinvoicePaperCommonReturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetBizErrorCode is BizErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoicePaperCommonReturnAPIRequest) SetBizErrorCode(_bizErrorCode string) error {
	r._bizErrorCode = _bizErrorCode
	r.Set("biz_error_code", _bizErrorCode)
	return nil
}

// GetBizErrorCode BizErrorCode Getter
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetBizErrorCode() string {
	return r._bizErrorCode
}

// SetBizErrorMsg is BizErrorMsg Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoicePaperCommonReturnAPIRequest) SetBizErrorMsg(_bizErrorMsg string) error {
	r._bizErrorMsg = _bizErrorMsg
	r.Set("biz_error_msg", _bizErrorMsg)
	return nil
}

// GetBizErrorMsg BizErrorMsg Getter
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetBizErrorMsg() string {
	return r._bizErrorMsg
}

// SetExtProps is ExtProps Setter
// 扩展信息
func (r *AlibabaEinvoicePaperCommonReturnAPIRequest) SetExtProps(_extProps string) error {
	r._extProps = _extProps
	r.Set("ext_props", _extProps)
	return nil
}

// GetExtProps ExtProps Getter
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetExtProps() string {
	return r._extProps
}

// SetSuccess is Success Setter
// 回传结果
func (r *AlibabaEinvoicePaperCommonReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaEinvoicePaperCommonReturnAPIRequest) GetSuccess() bool {
	return r._success
}
