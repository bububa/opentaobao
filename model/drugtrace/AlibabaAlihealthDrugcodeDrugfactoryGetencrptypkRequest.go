package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取加密公钥 API请求
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest struct {
    model.Params
    // 企业Id
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest{
    return &AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.drugfactory.getencrptypk"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest) GetRefEntId() string {
    return r._refEntId
}
