package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomemberhismemberdatagetprivyAPIRequest 会员历史备份数据查询 API请求
// taobao.member.hismemberdata.get.privy
//
// 会员历史备份数据分页查询，查询内容为等级，会员备份版本，会员nick等信息.
type TaobaomemberhismemberdatagetprivyAPIRequest struct {
	model.Params
	// 备份快照版本号（日期格式yyyyMMdd），可选参数，默认不填，则查找最近备份版本数据
	_backupDs string
	// 页面大小，必填，第几页
	_pageSize int64
	// 页码，必填，从1开始
	_currentPage int64
}

// NewTaobaomemberhismemberdatagetprivyRequest 初始化TaobaomemberhismemberdatagetprivyAPIRequest对象
func NewTaobaomemberhismemberdatagetprivyRequest() *TaobaomemberhismemberdatagetprivyAPIRequest {
	return &TaobaomemberhismemberdatagetprivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetApiMethodName() string {
	return "taobao.member.hismemberdata.get.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBackupDs is BackupDs Setter
// 备份快照版本号（日期格式yyyyMMdd），可选参数，默认不填，则查找最近备份版本数据
func (r *TaobaomemberhismemberdatagetprivyAPIRequest) SetBackupDs(_backupDs string) error {
	r._backupDs = _backupDs
	r.Set("backup_ds", _backupDs)
	return nil
}

// GetBackupDs BackupDs Getter
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetBackupDs() string {
	return r._backupDs
}

// SetPageSize is PageSize Setter
// 页面大小，必填，第几页
func (r *TaobaomemberhismemberdatagetprivyAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 页码，必填，从1开始
func (r *TaobaomemberhismemberdatagetprivyAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaomemberhismemberdatagetprivyAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
