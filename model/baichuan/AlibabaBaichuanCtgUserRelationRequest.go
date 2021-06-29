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
    app   string
    // 业务方的用户ID
    uid   string
    // 淘宝的用户ID
    tbUid   string
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
func (r *AlibabaBaichuanCtgUserRelationRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

// App Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetApp() string {
    return r.app
}
// Uid Setter
// 业务方的用户ID
func (r *AlibabaBaichuanCtgUserRelationRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetUid() string {
    return r.uid
}
// TbUid Setter
// 淘宝的用户ID
func (r *AlibabaBaichuanCtgUserRelationRequest) SetTbUid(tbUid string) error {
    r.tbUid = tbUid
    r.Set("tb_uid", tbUid)
    return nil
}

// TbUid Getter
func (r AlibabaBaichuanCtgUserRelationRequest) GetTbUid() string {
    return r.tbUid
}
