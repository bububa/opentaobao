package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印客户端监控信息收集 API请求
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集
*/
type CainiaoCloudprintClientinfoPutAPIRequest struct {
    model.Params
    // 客户端上传json数据
    _jsonData   string
}

// 初始化CainiaoCloudprintClientinfoPutAPIRequest对象
func NewCainiaoCloudprintClientinfoPutRequest() *CainiaoCloudprintClientinfoPutAPIRequest{
    return &CainiaoCloudprintClientinfoPutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.clientinfo.put"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JsonData Setter
// 客户端上传json数据
func (r *CainiaoCloudprintClientinfoPutAPIRequest) SetJsonData(_jsonData string) error {
    r._jsonData = _jsonData
    r.Set("json_data", _jsonData)
    return nil
}

// JsonData Getter
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetJsonData() string {
    return r._jsonData
}
