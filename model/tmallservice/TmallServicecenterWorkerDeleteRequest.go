package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除工人 API请求
tmall.servicecenter.worker.delete

删除工人信息。该接口为多个业务公用，部分字段可忽略。对于电器预约安装业务，同一个服务商，通过工人姓名+手机号+biz_type 保证唯一性。工人已存在才可以删除。
错误码如下
100000, 系统错误
100001, 工人信息校验失败
100002, 用户校验失败
100003, 操作失败
10004,工人信息为空
10005,服务商id为空或者服务商名称为空
10006, 工人不存在
10007, 工人已存在
10008, 缺少工人姓名
10009, 缺少工人电话
10010, 网点不存在
11000, category_id 无效
11001, biz_type 无效
20001,已查询到最后一页
*/
type TmallServicecenterWorkerDeleteAPIRequest struct {
    model.Params
    // 工人姓名
    _name   string
    // 工人手机号
    _phone   int64
    // 业务类型,电器预约安装业务填appliance_install
    _bizType   string
}

// 初始化TmallServicecenterWorkerDeleteAPIRequest对象
func NewTmallServicecenterWorkerDeleteRequest() *TmallServicecenterWorkerDeleteAPIRequest{
    return &TmallServicecenterWorkerDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerDeleteAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.delete"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 工人姓名
func (r *TmallServicecenterWorkerDeleteAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TmallServicecenterWorkerDeleteAPIRequest) GetName() string {
    return r._name
}
// Phone Setter
// 工人手机号
func (r *TmallServicecenterWorkerDeleteAPIRequest) SetPhone(_phone int64) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TmallServicecenterWorkerDeleteAPIRequest) GetPhone() int64 {
    return r._phone
}
// BizType Setter
// 业务类型,电器预约安装业务填appliance_install
func (r *TmallServicecenterWorkerDeleteAPIRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallServicecenterWorkerDeleteAPIRequest) GetBizType() string {
    return r._bizType
}
