package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码查询药品 APIRequest
alibaba.alihealth.drug.kyt.querydruginfo

通过追溯码查询药品信息
*/
type AlibabaAlihealthDrugKytQuerydruginfoRequest struct {
    model.Params

    // 码列表
    codeList   []string 

    // 物流企业refentid
    wuliuRefEntId   string 

    // 生产企业refentid
    huozhuRefEntId   string 

}

func NewAlibabaAlihealthDrugKytQuerydruginfoRequest() *AlibabaAlihealthDrugKytQuerydruginfoRequest{
    return &AlibabaAlihealthDrugKytQuerydruginfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querydruginfo"
}

func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetCodeList() []string {
    return r.codeList
}

func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetWuliuRefEntId(wuliuRefEntId string) error {
    r.wuliuRefEntId = wuliuRefEntId
    r.Set("wuliu_ref_ent_id", wuliuRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetWuliuRefEntId() string {
    return r.wuliuRefEntId
}

func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetHuozhuRefEntId(huozhuRefEntId string) error {
    r.huozhuRefEntId = huozhuRefEntId
    r.Set("huozhu_ref_ent_id", huozhuRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetHuozhuRefEntId() string {
    return r.huozhuRefEntId
}

