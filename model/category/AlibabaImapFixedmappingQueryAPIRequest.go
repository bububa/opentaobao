package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaImapFixedmappingQueryAPIRequest
查询两个渠道之间的固定映射关系，不通过算法兜底 API请求
alibaba.imap.fixedmapping.query

查询两个渠道之间的固定映射关系，不通过算法兜底 */
type AlibabaImapFixedmappingQueryAPIRequest struct {
	model.Params
	// 密码
	_password string
	// 账号
	_appName string
	// 源渠道ID
	_srcChannelId int64
	// 目标渠道ID列表
	_targetChannelIdList []int64
	// 目标渠道ID
	_targetCategoryId int64
	// 源类目ID
	_srcCategoryId int64
}

// NewAlibabaImapFixedmappingQueryRequest 初始化AlibabaImapFixedmappingQueryAPIRequest对象
func NewAlibabaImapFixedmappingQueryRequest() *AlibabaImapFixedmappingQueryAPIRequest {
	return &AlibabaImapFixedmappingQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaImapFixedmappingQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.imap.fixedmapping.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaImapFixedmappingQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Password Setter
// 密码
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetPassword(_password string) error {
	r._password = _password
	r.Set("password", _password)
	return nil
}

// Get Password Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetPassword() string {
	return r._password
}

// Set is AppName Setter
// 账号
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// Get AppName Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetAppName() string {
	return r._appName
}

// Set is SrcChannelId Setter
// 源渠道ID
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetSrcChannelId(_srcChannelId int64) error {
	r._srcChannelId = _srcChannelId
	r.Set("src_channel_id", _srcChannelId)
	return nil
}

// Get SrcChannelId Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetSrcChannelId() int64 {
	return r._srcChannelId
}

// Set is TargetChannelIdList Setter
// 目标渠道ID列表
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetTargetChannelIdList(_targetChannelIdList []int64) error {
	r._targetChannelIdList = _targetChannelIdList
	r.Set("target_channel_id_list", _targetChannelIdList)
	return nil
}

// Get TargetChannelIdList Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetTargetChannelIdList() []int64 {
	return r._targetChannelIdList
}

// Set is TargetCategoryId Setter
// 目标渠道ID
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetTargetCategoryId(_targetCategoryId int64) error {
	r._targetCategoryId = _targetCategoryId
	r.Set("target_category_id", _targetCategoryId)
	return nil
}

// Get TargetCategoryId Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetTargetCategoryId() int64 {
	return r._targetCategoryId
}

// Set is SrcCategoryId Setter
// 源类目ID
func (r *AlibabaImapFixedmappingQueryAPIRequest) SetSrcCategoryId(_srcCategoryId int64) error {
	r._srcCategoryId = _srcCategoryId
	r.Set("src_category_id", _srcCategoryId)
	return nil
}

// Get SrcCategoryId Getter
func (r AlibabaImapFixedmappingQueryAPIRequest) GetSrcCategoryId() int64 {
	return r._srcCategoryId
}
