package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest 运营位管理-联盟一体机下线运营位内容 API请求
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
type YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest struct {
	model.Params
	// 联盟：TV_OTT,一体机：TV_ALLINONE
	_terminalType string
	// 元数据主键id
	_id int64
}

// NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest 初始化YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest对象
func NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest() *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest {
	return &YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) Reset() {
	r._terminalType = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.offlinelauncheritem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 联盟：TV_OTT,一体机：TV_ALLINONE
func (r *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetId is Id Setter
// 元数据主键id
func (r *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetId() int64 {
	return r._id
}

var poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest()
	},
}

// GetYunosTvpubadminContentTableauditOfflinelauncheritemRequest 从 sync.Pool 获取 YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest
func GetYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest() *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest {
	return poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest.Get().(*YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest)
}

// ReleaseYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest 将 YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest(v *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest.Put(v)
}
