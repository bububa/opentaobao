package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新维修进度 API请求
tmall.servicecenter.workcard.repairprogress.update

提供给外部合作服务商的维修进度更改接口
*/
type TmallServicecenterWorkcardRepairprogressUpdateRequest struct {
    model.Params
    // 图片列表
    _picUrlList   []string
    // 请求节点的动作描述，唯一标识一个节点
    _action   string
    // 工单id
    _workcardId   int64
    // 真实接单服务商账号Nick
    _realTpNick   string
    // 服务目标物瑕疵信息
    _targetGoodsDefects   string
    // 衣服，鞋子
    _receivedGoods   string
}

// 初始化TmallServicecenterWorkcardRepairprogressUpdateRequest对象
func NewTmallServicecenterWorkcardRepairprogressUpdateRequest() *TmallServicecenterWorkcardRepairprogressUpdateRequest{
    return &TmallServicecenterWorkcardRepairprogressUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.repairprogress.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PicUrlList Setter
// 图片列表
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetPicUrlList(_picUrlList []string) error {
    r._picUrlList = _picUrlList
    r.Set("pic_url_list", _picUrlList)
    return nil
}

// PicUrlList Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetPicUrlList() []string {
    return r._picUrlList
}
// Action Setter
// 请求节点的动作描述，唯一标识一个节点
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetAction() string {
    return r._action
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// RealTpNick Setter
// 真实接单服务商账号Nick
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetRealTpNick(_realTpNick string) error {
    r._realTpNick = _realTpNick
    r.Set("real_tp_nick", _realTpNick)
    return nil
}

// RealTpNick Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetRealTpNick() string {
    return r._realTpNick
}
// TargetGoodsDefects Setter
// 服务目标物瑕疵信息
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetTargetGoodsDefects(_targetGoodsDefects string) error {
    r._targetGoodsDefects = _targetGoodsDefects
    r.Set("target_goods_defects", _targetGoodsDefects)
    return nil
}

// TargetGoodsDefects Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetTargetGoodsDefects() string {
    return r._targetGoodsDefects
}
// ReceivedGoods Setter
// 衣服，鞋子
func (r *TmallServicecenterWorkcardRepairprogressUpdateRequest) SetReceivedGoods(_receivedGoods string) error {
    r._receivedGoods = _receivedGoods
    r.Set("received_goods", _receivedGoods)
    return nil
}

// ReceivedGoods Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateRequest) GetReceivedGoods() string {
    return r._receivedGoods
}
