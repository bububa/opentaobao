package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发出奇门事件 API请求
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。
*/
type TaobaoQimenEventProduceRequest struct {
    model.Params
    // 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
    _status   string
    // 淘宝订单号
    _tid   string
    // JSON格式扩展字段
    _ext   string
    // 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
    _platform   string
    // 订单创建时间,数字
    _create   int64
    // 外部商家名称。必须同时填写platform
    _nick   string
}

// 初始化TaobaoQimenEventProduceRequest对象
func NewTaobaoQimenEventProduceRequest() *TaobaoQimenEventProduceRequest{
    return &TaobaoQimenEventProduceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEventProduceRequest) GetApiMethodName() string {
    return "taobao.qimen.event.produce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEventProduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
func (r *TaobaoQimenEventProduceRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoQimenEventProduceRequest) GetStatus() string {
    return r._status
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoQimenEventProduceRequest) SetTid(_tid string) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoQimenEventProduceRequest) GetTid() string {
    return r._tid
}
// Ext Setter
// JSON格式扩展字段
func (r *TaobaoQimenEventProduceRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoQimenEventProduceRequest) GetExt() string {
    return r._ext
}
// Platform Setter
// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
func (r *TaobaoQimenEventProduceRequest) SetPlatform(_platform string) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r TaobaoQimenEventProduceRequest) GetPlatform() string {
    return r._platform
}
// Create Setter
// 订单创建时间,数字
func (r *TaobaoQimenEventProduceRequest) SetCreate(_create int64) error {
    r._create = _create
    r.Set("create", _create)
    return nil
}

// Create Getter
func (r TaobaoQimenEventProduceRequest) GetCreate() int64 {
    return r._create
}
// Nick Setter
// 外部商家名称。必须同时填写platform
func (r *TaobaoQimenEventProduceRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoQimenEventProduceRequest) GetNick() string {
    return r._nick
}
