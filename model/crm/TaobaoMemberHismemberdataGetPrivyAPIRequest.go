package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMemberHismemberdataGetPrivyAPIRequest 会员历史备份数据查询 API请求
// taobao.member.hismemberdata.get.privy
//
// 会员历史备份数据分页查询，查询内容为等级，会员备份版本，会员nick等信息.
type TaobaoMemberHismemberdataGetPrivyAPIRequest struct {
	model.Params
	// 备份快照版本号（日期格式yyyyMMdd），可选参数，默认不填，则查找最近备份版本数据
	_backupDs string
	// 页面大小，必填，第几页
	_pageSize int64
	// 页码，必填，从1开始
	_currentPage int64
}

// NewTaobaoMemberHismemberdataGetPrivyRequest 初始化TaobaoMemberHismemberdataGetPrivyAPIRequest对象
func NewTaobaoMemberHismemberdataGetPrivyRequest() *TaobaoMemberHismemberdataGetPrivyAPIRequest {
	return &TaobaoMemberHismemberdataGetPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMemberHismemberdataGetPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.member.hismemberdata.get.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMemberHismemberdataGetPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBackupDs is BackupDs Setter
// 备份快照版本号（日期格式yyyyMMdd），可选参数，默认不填，则查找最近备份版本数据
func (r *TaobaoMemberHismemberdataGetPrivyAPIRequest) SetBackupDs(_backupDs string) error {
	r._backupDs = _backupDs
	r.Set("backup_ds", _backupDs)
	return nil
}

// GetBackupDs BackupDs Getter
func (r TaobaoMemberHismemberdataGetPrivyAPIRequest) GetBackupDs() string {
	return r._backupDs
}

// SetPageSize is PageSize Setter
// 页面大小，必填，第几页
func (r *TaobaoMemberHismemberdataGetPrivyAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoMemberHismemberdataGetPrivyAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 页码，必填，从1开始
func (r *TaobaoMemberHismemberdataGetPrivyAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoMemberHismemberdataGetPrivyAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
