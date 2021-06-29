package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
影视桌面运营通用接口 APIRequest
yunos.tvpubadmin.epg.desktop.operation

影视桌面运营通用接口
*/
type YunosTvpubadminEpgDesktopOperationRequest struct {
    model.Params

    // 操作对象实体
    entityType   string 

    // 操作类型
    actionType   string 

    // 具体入参
    parameter   string 

}

func NewYunosTvpubadminEpgDesktopOperationRequest() *YunosTvpubadminEpgDesktopOperationRequest{
    return &YunosTvpubadminEpgDesktopOperationRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminEpgDesktopOperationRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.epg.desktop.operation"
}

func (r YunosTvpubadminEpgDesktopOperationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminEpgDesktopOperationRequest) SetEntityType(entityType string) error {
    r.entityType = entityType
    r.Set("entity_type", entityType)
    return nil
}

func (r YunosTvpubadminEpgDesktopOperationRequest) GetEntityType() string {
    return r.entityType
}

func (r *YunosTvpubadminEpgDesktopOperationRequest) SetActionType(actionType string) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

func (r YunosTvpubadminEpgDesktopOperationRequest) GetActionType() string {
    return r.actionType
}

func (r *YunosTvpubadminEpgDesktopOperationRequest) SetParameter(parameter string) error {
    r.parameter = parameter
    r.Set("parameter", parameter)
    return nil
}

func (r YunosTvpubadminEpgDesktopOperationRequest) GetParameter() string {
    return r.parameter
}

