package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码的药品信息查询 API请求
alibaba.alihealth.drug.kyt.getcodebaseinfo

提供根据码查询码基本信息接口
*/
type AlibabaAlihealthDrugKytGetcodebaseinfoRequest struct {
    model.Params
    // 码
    _code   string
    // 企业唯一标识
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugKytGetcodebaseinfoRequest对象
func NewAlibabaAlihealthDrugKytGetcodebaseinfoRequest() *AlibabaAlihealthDrugKytGetcodebaseinfoRequest{
    return &AlibabaAlihealthDrugKytGetcodebaseinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getcodebaseinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 码
func (r *AlibabaAlihealthDrugKytGetcodebaseinfoRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytGetcodebaseinfoRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetRefEntId() string {
    return r._refEntId
}
