package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest 运营位管理-联盟一体机下线运营位内容 API请求
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
type YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest struct {
	model.Params
	// 元数据主键id
	_id int64
	// 联盟：TV_OTT,一体机：TV_ALLINONE
	_terminalType string
}

// NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest 初始化YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest对象
func NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest() *YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest {
	return &YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.offlinelauncheritem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditOfflinelauncheritemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
