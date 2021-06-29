package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品一品多码维护操作 APIRequest
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作
*/
type AlibabaWdkItemMorebarcodeOpsRequest struct {
    model.Params

    // bean
    updateMoreBarCodeRequestBean   *UpdateMoreBarCodeRequestBean 

}

func NewAlibabaWdkItemMorebarcodeOpsRequest() *AlibabaWdkItemMorebarcodeOpsRequest{
    return &AlibabaWdkItemMorebarcodeOpsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMorebarcodeOpsRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.morebarcode.ops"
}

func (r AlibabaWdkItemMorebarcodeOpsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMorebarcodeOpsRequest) SetUpdateMoreBarCodeRequestBean(updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean) error {
    r.updateMoreBarCodeRequestBean = updateMoreBarCodeRequestBean
    r.Set("update_more_bar_code_request_bean", updateMoreBarCodeRequestBean)
    return nil
}

func (r AlibabaWdkItemMorebarcodeOpsRequest) GetUpdateMoreBarCodeRequestBean() *UpdateMoreBarCodeRequestBean {
    return r.updateMoreBarCodeRequestBean
}

