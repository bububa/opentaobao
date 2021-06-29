package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收盲底文件删除日志 API请求
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

// 初始化AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.blindfiledellog"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 药厂企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetRefEntId() string {
    return r.refEntId
}
// Operator Setter
// 操作人
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetOperator() string {
    return r.operator
}
// BlindFileDeleteTime Setter
// 盲底文件删除时间
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetBlindFileDeleteTime(blindFileDeleteTime string) error {
    r.blindFileDeleteTime = blindFileDeleteTime
    r.Set("blind_file_delete_time", blindFileDeleteTime)
    return nil
}

// BlindFileDeleteTime Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetBlindFileDeleteTime() string {
    return r.blindFileDeleteTime
}
// BlindFileNames Setter
// 盲底文件名称，多个盲底文件用,分隔
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) SetBlindFileNames(blindFileNames string) error {
    r.blindFileNames = blindFileNames
    r.Set("blind_file_names", blindFileNames)
    return nil
}

// BlindFileNames Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest) GetBlindFileNames() string {
    return r.blindFileNames
}
