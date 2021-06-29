package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票通用回传接口 API请求
alibaba.einvoice.paper.common.return

纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
*/
type AlibabaEinvoicePaperCommonReturnRequest struct {
    model.Params
    // 请求索引
    _reqIndex   string
    // 回传结果
    _success   bool
    // 错误码，success=false时必填
    _bizErrorCode   string
    // 错误信息，success=false时必填
    _bizErrorMsg   string
    // 扩展信息
    _extProps   string
}

// 初始化AlibabaEinvoicePaperCommonReturnRequest对象
func NewAlibabaEinvoicePaperCommonReturnRequest() *AlibabaEinvoicePaperCommonReturnRequest{
    return &AlibabaEinvoicePaperCommonReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePaperCommonReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.common.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePaperCommonReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReqIndex Setter
// 请求索引
func (r *AlibabaEinvoicePaperCommonReturnRequest) SetReqIndex(_reqIndex string) error {
    r._reqIndex = _reqIndex
    r.Set("req_index", _reqIndex)
    return nil
}

// ReqIndex Getter
func (r AlibabaEinvoicePaperCommonReturnRequest) GetReqIndex() string {
    return r._reqIndex
}
// Success Setter
// 回传结果
func (r *AlibabaEinvoicePaperCommonReturnRequest) SetSuccess(_success bool) error {
    r._success = _success
    r.Set("success", _success)
    return nil
}

// Success Getter
func (r AlibabaEinvoicePaperCommonReturnRequest) GetSuccess() bool {
    return r._success
}
// BizErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoicePaperCommonReturnRequest) SetBizErrorCode(_bizErrorCode string) error {
    r._bizErrorCode = _bizErrorCode
    r.Set("biz_error_code", _bizErrorCode)
    return nil
}

// BizErrorCode Getter
func (r AlibabaEinvoicePaperCommonReturnRequest) GetBizErrorCode() string {
    return r._bizErrorCode
}
// BizErrorMsg Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoicePaperCommonReturnRequest) SetBizErrorMsg(_bizErrorMsg string) error {
    r._bizErrorMsg = _bizErrorMsg
    r.Set("biz_error_msg", _bizErrorMsg)
    return nil
}

// BizErrorMsg Getter
func (r AlibabaEinvoicePaperCommonReturnRequest) GetBizErrorMsg() string {
    return r._bizErrorMsg
}
// ExtProps Setter
// 扩展信息
func (r *AlibabaEinvoicePaperCommonReturnRequest) SetExtProps(_extProps string) error {
    r._extProps = _extProps
    r.Set("ext_props", _extProps)
    return nil
}

// ExtProps Getter
func (r AlibabaEinvoicePaperCommonReturnRequest) GetExtProps() string {
    return r._extProps
}
