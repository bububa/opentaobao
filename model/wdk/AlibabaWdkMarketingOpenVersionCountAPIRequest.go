package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenVersionCountAPIRequest 版本数量查询 API请求
// alibaba.wdk.marketing.open.version.count
//
// 版本数量查询
type AlibabaWdkMarketingOpenVersionCountAPIRequest struct {
	model.Params
	// 查询版本号
	_versionId int64
	// 操作Id
	_operateId string
}

// NewAlibabaWdkMarketingOpenVersionCountRequest 初始化AlibabaWdkMarketingOpenVersionCountAPIRequest对象
func NewAlibabaWdkMarketingOpenVersionCountRequest() *AlibabaWdkMarketingOpenVersionCountAPIRequest {
	return &AlibabaWdkMarketingOpenVersionCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenVersionCountAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.version.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenVersionCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVersionId is VersionId Setter
// 查询版本号
func (r *AlibabaWdkMarketingOpenVersionCountAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r AlibabaWdkMarketingOpenVersionCountAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetOperateId is OperateId Setter
// 操作Id
func (r *AlibabaWdkMarketingOpenVersionCountAPIRequest) SetOperateId(_operateId string) error {
	r._operateId = _operateId
	r.Set("operate_id", _operateId)
	return nil
}

// GetOperateId OperateId Getter
func (r AlibabaWdkMarketingOpenVersionCountAPIRequest) GetOperateId() string {
	return r._operateId
}
