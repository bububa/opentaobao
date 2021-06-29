package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管理-联盟一体机下线运营位内容 API请求
yunos.tvpubadmin.content.tableaudit.offlinelauncheritem

运营位管理-联盟一体机下线运营位内容
*/
type YunosTvpubadminContentTableauditOfflinelauncheritemRequest struct {
    model.Params
    // 元数据主键id
    _id   int64
    // 联盟：TV_OTT,一体机：TV_ALLINONE
    _terminalType   string
}

// 初始化YunosTvpubadminContentTableauditOfflinelauncheritemRequest对象
func NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest() *YunosTvpubadminContentTableauditOfflinelauncheritemRequest{
    return &YunosTvpubadminContentTableauditOfflinelauncheritemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.offlinelauncheritem"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 元数据主键id
func (r *YunosTvpubadminContentTableauditOfflinelauncheritemRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetId() int64 {
    return r._id
}
// TerminalType Setter
// 联盟：TV_OTT,一体机：TV_ALLINONE
func (r *YunosTvpubadminContentTableauditOfflinelauncheritemRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetTerminalType() string {
    return r._terminalType
}
