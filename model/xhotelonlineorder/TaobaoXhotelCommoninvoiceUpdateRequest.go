package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息更新接口 API请求
taobao.xhotel.commoninvoice.update

常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
*/
type TaobaoXhotelCommoninvoiceUpdateRequest struct {
    model.Params
    // 无
    _commonInvoiceInfoParam   *CommonInvoiceInfo
}

// 初始化TaobaoXhotelCommoninvoiceUpdateRequest对象
func NewTaobaoXhotelCommoninvoiceUpdateRequest() *TaobaoXhotelCommoninvoiceUpdateRequest{
    return &TaobaoXhotelCommoninvoiceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CommonInvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelCommoninvoiceUpdateRequest) SetCommonInvoiceInfoParam(_commonInvoiceInfoParam *CommonInvoiceInfo) error {
    r._commonInvoiceInfoParam = _commonInvoiceInfoParam
    r.Set("common_invoice_info_param", _commonInvoiceInfoParam)
    return nil
}

// CommonInvoiceInfoParam Getter
func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetCommonInvoiceInfoParam() *CommonInvoiceInfo {
    return r._commonInvoiceInfoParam
}
