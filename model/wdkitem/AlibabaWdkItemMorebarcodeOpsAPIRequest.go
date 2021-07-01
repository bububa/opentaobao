package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品一品多码维护操作 API请求
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作
*/
type AlibabaWdkItemMorebarcodeOpsAPIRequest struct {
    model.Params
    // bean
    _updateMoreBarCodeRequestBean   *UpdateMoreBarCodeRequestBean
}

// 初始化AlibabaWdkItemMorebarcodeOpsAPIRequest对象
func NewAlibabaWdkItemMorebarcodeOpsRequest() *AlibabaWdkItemMorebarcodeOpsAPIRequest{
    return &AlibabaWdkItemMorebarcodeOpsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.morebarcode.ops"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdateMoreBarCodeRequestBean Setter
// bean
func (r *AlibabaWdkItemMorebarcodeOpsAPIRequest) SetUpdateMoreBarCodeRequestBean(_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean) error {
    r._updateMoreBarCodeRequestBean = _updateMoreBarCodeRequestBean
    r.Set("update_more_bar_code_request_bean", _updateMoreBarCodeRequestBean)
    return nil
}

// UpdateMoreBarCodeRequestBean Getter
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetUpdateMoreBarCodeRequestBean() *UpdateMoreBarCodeRequestBean {
    return r._updateMoreBarCodeRequestBean
}
