package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管理-联盟一体机下线运营位内容 APIRequest
yunos.tvpubadmin.content.tableaudit.offlinelauncheritem

运营位管理-联盟一体机下线运营位内容
*/
type YunosTvpubadminContentTableauditOfflinelauncheritemRequest struct {
    model.Params

    // 元数据主键id
    id   int64 

    // 联盟：TV_OTT,一体机：TV_ALLINONE
    terminalType   string 

}

func NewYunosTvpubadminContentTableauditOfflinelauncheritemRequest() *YunosTvpubadminContentTableauditOfflinelauncheritemRequest{
    return &YunosTvpubadminContentTableauditOfflinelauncheritemRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.offlinelauncheritem"
}

func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTableauditOfflinelauncheritemRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetId() int64 {
    return r.id
}

func (r *YunosTvpubadminContentTableauditOfflinelauncheritemRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r YunosTvpubadminContentTableauditOfflinelauncheritemRequest) GetTerminalType() string {
    return r.terminalType
}

