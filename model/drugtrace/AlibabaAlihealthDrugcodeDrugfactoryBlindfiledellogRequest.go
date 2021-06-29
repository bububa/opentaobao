package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收盲底文件删除日志 APIRequest
alibaba.alihealth.drugcode.drugfactory.blindfiledellog

临床用药试验-接收盲底文件删除日志
*/
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest struct {
    model.Params

    // 药厂企业id
    refEntId   string 

    // 操作人
    operator   string 

    // 盲底文件删除时间
    blindFileDeleteTime   string 

    // 盲底文件名称，多个盲底文件用,分隔
    blindFileNames   string 

}

func NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.blindfiledellog"
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetOperator() string {
    return r.operator
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetBlindFileDeleteTime(blindFileDeleteTime string) error {
    r.blindFileDeleteTime = blindFileDeleteTime
    r.Set("blind_file_delete_time", blindFileDeleteTime)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetBlindFileDeleteTime() string {
    return r.blindFileDeleteTime
}

func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetBlindFileNames(blindFileNames string) error {
    r.blindFileNames = blindFileNames
    r.Set("blind_file_names", blindFileNames)
    return nil
}

func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetBlindFileNames() string {
    return r.blindFileNames
}

