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
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest struct {
    model.Params
    // 药厂企业id
    _refEntId   string
    // 操作人
    _operator   string
    // 盲底文件删除时间
    _blindFileDeleteTime   string
    // 盲底文件名称，多个盲底文件用,分隔
    _blindFileNames   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.blindfiledellog"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 药厂企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// Operator Setter
// 操作人
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetOperator() string {
    return r._operator
}
// BlindFileDeleteTime Setter
// 盲底文件删除时间
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetBlindFileDeleteTime(_blindFileDeleteTime string) error {
    r._blindFileDeleteTime = _blindFileDeleteTime
    r.Set("blind_file_delete_time", _blindFileDeleteTime)
    return nil
}

// BlindFileDeleteTime Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetBlindFileDeleteTime() string {
    return r._blindFileDeleteTime
}
// BlindFileNames Setter
// 盲底文件名称，多个盲底文件用,分隔
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetBlindFileNames(_blindFileNames string) error {
    r._blindFileNames = _blindFileNames
    r.Set("blind_file_names", _blindFileNames)
    return nil
}

// BlindFileNames Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetBlindFileNames() string {
    return r._blindFileNames
}
