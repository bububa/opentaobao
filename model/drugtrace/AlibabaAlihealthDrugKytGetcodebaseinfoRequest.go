package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码的药品信息查询 APIRequest
alibaba.alihealth.drug.kyt.getcodebaseinfo

提供根据码查询码基本信息接口
*/
type AlibabaAlihealthDrugKytGetcodebaseinfoRequest struct {
    model.Params

    // 码
    code   string 

    // 企业唯一标识
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytGetcodebaseinfoRequest() *AlibabaAlihealthDrugKytGetcodebaseinfoRequest{
    return &AlibabaAlihealthDrugKytGetcodebaseinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getcodebaseinfo"
}

func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetcodebaseinfoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugKytGetcodebaseinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetcodebaseinfoRequest) GetRefEntId() string {
    return r.refEntId
}

