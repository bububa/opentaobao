package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码查询药品 API请求
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

// 初始化AlibabaAlihealthDrugKytQuerydruginfoRequest对象
func NewAlibabaAlihealthDrugKytQuerydruginfoRequest() *AlibabaAlihealthDrugKytQuerydruginfoRequest{
    return &AlibabaAlihealthDrugKytQuerydruginfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querydruginfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeList Setter
// 码列表
func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetCodeList() []string {
    return r.codeList
}
// WuliuRefEntId Setter
// 物流企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetWuliuRefEntId(wuliuRefEntId string) error {
    r.wuliuRefEntId = wuliuRefEntId
    r.Set("wuliu_ref_ent_id", wuliuRefEntId)
    return nil
}

// WuliuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetWuliuRefEntId() string {
    return r.wuliuRefEntId
}
// HuozhuRefEntId Setter
// 生产企业refentid
func (r *AlibabaAlihealthDrugKytQuerydruginfoRequest) SetHuozhuRefEntId(huozhuRefEntId string) error {
    r.huozhuRefEntId = huozhuRefEntId
    r.Set("huozhu_ref_ent_id", huozhuRefEntId)
    return nil
}

// HuozhuRefEntId Getter
func (r AlibabaAlihealthDrugKytQuerydruginfoRequest) GetHuozhuRefEntId() string {
    return r.huozhuRefEntId
}
