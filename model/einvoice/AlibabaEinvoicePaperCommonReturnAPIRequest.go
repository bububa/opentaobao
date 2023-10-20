package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicepapercommonreturnAPIRequest 纸票通用回传接口 API请求
// alibaba.einvoice.paper.common.return
//
// 纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
type AlibabaeinvoicepapercommonreturnAPIRequest struct {
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

// NewAlibabaeinvoicepapercommonreturnRequest 初始化AlibabaeinvoicepapercommonreturnAPIRequest对象
func NewAlibabaeinvoicepapercommonreturnRequest() *AlibabaeinvoicepapercommonreturnAPIRequest {
	return &AlibabaeinvoicepapercommonreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.paper.common.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqIndex is ReqIndex Setter
// 请求索引
func (r *AlibabaeinvoicepapercommonreturnAPIRequest) SetReqIndex(_reqIndex string) error {
	r._reqIndex = _reqIndex
	r.Set("req_index", _reqIndex)
	return nil
}

// GetReqIndex ReqIndex Getter
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetReqIndex() string {
	return r._reqIndex
}

// SetBizErrorCode is BizErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaeinvoicepapercommonreturnAPIRequest) SetBizErrorCode(_bizErrorCode string) error {
	r._bizErrorCode = _bizErrorCode
	r.Set("biz_error_code", _bizErrorCode)
	return nil
}

// GetBizErrorCode BizErrorCode Getter
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetBizErrorCode() string {
	return r._bizErrorCode
}

// SetBizErrorMsg is BizErrorMsg Setter
// 错误信息，success=false时必填
func (r *AlibabaeinvoicepapercommonreturnAPIRequest) SetBizErrorMsg(_bizErrorMsg string) error {
	r._bizErrorMsg = _bizErrorMsg
	r.Set("biz_error_msg", _bizErrorMsg)
	return nil
}

// GetBizErrorMsg BizErrorMsg Getter
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetBizErrorMsg() string {
	return r._bizErrorMsg
}

// SetExtProps is ExtProps Setter
// 扩展信息
func (r *AlibabaeinvoicepapercommonreturnAPIRequest) SetExtProps(_extProps string) error {
	r._extProps = _extProps
	r.Set("ext_props", _extProps)
	return nil
}

// GetExtProps ExtProps Getter
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetExtProps() string {
	return r._extProps
}

// SetSuccess is Success Setter
// 回传结果
func (r *AlibabaeinvoicepapercommonreturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaeinvoicepapercommonreturnAPIRequest) GetSuccess() bool {
	return r._success
}
