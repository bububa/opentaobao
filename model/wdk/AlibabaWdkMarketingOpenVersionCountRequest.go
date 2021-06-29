package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
版本数量查询 API请求
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
type AlibabaWdkMarketingOpenVersionCountRequest struct {
    model.Params
    // 查询版本号
    _versionId   int64
    // 操作Id
    _operateId   string
}

// 初始化AlibabaWdkMarketingOpenVersionCountRequest对象
func NewAlibabaWdkMarketingOpenVersionCountRequest() *AlibabaWdkMarketingOpenVersionCountRequest{
    return &AlibabaWdkMarketingOpenVersionCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenVersionCountRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.open.version.count"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenVersionCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VersionId Setter
// 查询版本号
func (r *AlibabaWdkMarketingOpenVersionCountRequest) SetVersionId(_versionId int64) error {
    r._versionId = _versionId
    r.Set("version_id", _versionId)
    return nil
}

// VersionId Getter
func (r AlibabaWdkMarketingOpenVersionCountRequest) GetVersionId() int64 {
    return r._versionId
}
// OperateId Setter
// 操作Id
func (r *AlibabaWdkMarketingOpenVersionCountRequest) SetOperateId(_operateId string) error {
    r._operateId = _operateId
    r.Set("operate_id", _operateId)
    return nil
}

// OperateId Getter
func (r AlibabaWdkMarketingOpenVersionCountRequest) GetOperateId() string {
    return r._operateId
}
