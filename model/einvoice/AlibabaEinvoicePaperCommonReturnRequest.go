package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
纸票通用回传接口 APIRequest
alibaba.einvoice.paper.common.return

纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
*/
type AlibabaEinvoicePaperCommonReturnRequest struct {
    model.Params

    // 请求索引
    reqIndex   string 

    // 回传结果
    success   bool 

    // 错误码，success=false时必填
    bizErrorCode   string 

    // 错误信息，success=false时必填
    bizErrorMsg   string 

    // 扩展信息
    extProps   string 

}

func NewAlibabaEinvoicePaperCommonReturnRequest() *AlibabaEinvoicePaperCommonReturnRequest{
    return &AlibabaEinvoicePaperCommonReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.paper.common.return"
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoicePaperCommonReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

func (r *AlibabaEinvoicePaperCommonReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetSuccess() bool {
    return r.success
}

func (r *AlibabaEinvoicePaperCommonReturnRequest) SetBizErrorCode(bizErrorCode string) error {
    r.bizErrorCode = bizErrorCode
    r.Set("biz_error_code", bizErrorCode)
    return nil
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetBizErrorCode() string {
    return r.bizErrorCode
}

func (r *AlibabaEinvoicePaperCommonReturnRequest) SetBizErrorMsg(bizErrorMsg string) error {
    r.bizErrorMsg = bizErrorMsg
    r.Set("biz_error_msg", bizErrorMsg)
    return nil
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetBizErrorMsg() string {
    return r.bizErrorMsg
}

func (r *AlibabaEinvoicePaperCommonReturnRequest) SetExtProps(extProps string) error {
    r.extProps = extProps
    r.Set("ext_props", extProps)
    return nil
}

func (r AlibabaEinvoicePaperCommonReturnRequest) GetExtProps() string {
    return r.extProps
}

