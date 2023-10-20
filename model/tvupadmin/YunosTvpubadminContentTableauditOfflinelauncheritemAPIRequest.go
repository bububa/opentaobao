package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest 运营位管理-联盟一体机下线运营位内容 API请求
// yunos.tvpubadmin.content.tableaudit.offlinelauncheritem
//
// 运营位管理-联盟一体机下线运营位内容
type YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest struct {
	model.Params
	// 联盟：TV_OTT,一体机：TV_ALLINONE
	_terminalType string
	// 元数据主键id
	_id int64
}

// NewYunostvpubadmincontenttableauditofflinelauncheritemRequest 初始化YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest对象
func NewYunostvpubadmincontenttableauditofflinelauncheritemRequest() *YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest {
	return &YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.offlinelauncheritem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 联盟：TV_OTT,一体机：TV_ALLINONE
func (r *YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetId is Id Setter
// 元数据主键id
func (r *YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmincontenttableauditofflinelauncheritemAPIRequest) GetId() int64 {
	return r._id
}
