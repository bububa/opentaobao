package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopenversioncountAPIRequest 版本数量查询 API请求
// alibaba.wdk.marketing.open.version.count
//
// 版本数量查询
type AlibabawdkmarketingopenversioncountAPIRequest struct {
	model.Params
	// 操作Id
	_operateId string
	// 查询版本号
	_versionId int64
}

// NewAlibabawdkmarketingopenversioncountRequest 初始化AlibabawdkmarketingopenversioncountAPIRequest对象
func NewAlibabawdkmarketingopenversioncountRequest() *AlibabawdkmarketingopenversioncountAPIRequest {
	return &AlibabawdkmarketingopenversioncountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopenversioncountAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.version.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopenversioncountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopenversioncountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperateId is OperateId Setter
// 操作Id
func (r *AlibabawdkmarketingopenversioncountAPIRequest) SetOperateId(_operateId string) error {
	r._operateId = _operateId
	r.Set("operate_id", _operateId)
	return nil
}

// GetOperateId OperateId Getter
func (r AlibabawdkmarketingopenversioncountAPIRequest) GetOperateId() string {
	return r._operateId
}

// SetVersionId is VersionId Setter
// 查询版本号
func (r *AlibabawdkmarketingopenversioncountAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r AlibabawdkmarketingopenversioncountAPIRequest) GetVersionId() int64 {
	return r._versionId
}
