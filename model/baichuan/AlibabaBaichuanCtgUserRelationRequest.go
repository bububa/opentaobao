package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户 API请求
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系
*/
type AlibabaBaichuanCtgUserRelationRequest struct {
    model.Params
    // 调用的业务方
    _app   string
    // 业务方的用户ID
    _uid   string
    // 淘宝的用户ID
    _tbUid   string
}

// 初始化AlibabaBaichuanCtgUserRelationRequest对象
func NewAlibabaBaichuanCtgUserRelationRequest() *AlibabaBaichuanCtgUserRelationRequest{
    return &AlibabaBaichuanCtgUserRelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgUserRelationRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.user.relation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgUserRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// App Setter
// 调用的业务方
func (r *AlibabaBaichuanCtgUserRelationRequest) SetApp(_app string) error {
    r._app = _app
    r.Set("app", _app)
    return nil
}

// App Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetApp() string {
    return r._app
}
// Uid Setter
// 业务方的用户ID
func (r *AlibabaBaichuanCtgUserRelationRequest) SetUid(_uid string) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetUid() string {
    return r._uid
}
// TbUid Setter
// 淘宝的用户ID
func (r *AlibabaBaichuanCtgUserRelationRequest) SetTbUid(_tbUid string) error {
    r._tbUid = _tbUid
    r.Set("tb_uid", _tbUid)
    return nil
}

// TbUid Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetTbUid() string {
    return r._tbUid
}
