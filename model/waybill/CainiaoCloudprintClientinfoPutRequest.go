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
type CainiaoCloudprintClientinfoPutRequest struct {
    model.Params
    // 客户端上传json数据
    _jsonData   string
}

// 初始化CainiaoCloudprintClientinfoPutRequest对象
func NewCainiaoCloudprintClientinfoPutRequest() *CainiaoCloudprintClientinfoPutRequest{
    return &CainiaoCloudprintClientinfoPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintClientinfoPutRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.clientinfo.put"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintClientinfoPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JsonData Setter
// 客户端上传json数据
func (r *CainiaoCloudprintClientinfoPutRequest) SetJsonData(_jsonData string) error {
    r._jsonData = _jsonData
    r.Set("json_data", _jsonData)
    return nil
}

// JsonData Getter
func (r CainiaoCloudprintClientinfoPutRequest) GetJsonData() string {
    return r._jsonData
}
