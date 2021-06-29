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
    status   string
    // 淘宝订单号
    tid   string
    // JSON格式扩展字段
    ext   string
    // 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
    platform   string
    // 订单创建时间,数字
    create   int64
    // 外部商家名称。必须同时填写platform
    nick   string
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
func (r *TaobaoQimenEventProduceRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoQimenEventProduceRequest) GetStatus() string {
    return r.status
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoQimenEventProduceRequest) SetTid(tid string) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoQimenEventProduceRequest) GetTid() string {
    return r.tid
}
// Ext Setter
// JSON格式扩展字段
func (r *TaobaoQimenEventProduceRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoQimenEventProduceRequest) GetExt() string {
    return r.ext
}
// Platform Setter
// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
func (r *TaobaoQimenEventProduceRequest) SetPlatform(platform string) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r TaobaoQimenEventProduceRequest) GetPlatform() string {
    return r.platform
}
// Create Setter
// 订单创建时间,数字
func (r *TaobaoQimenEventProduceRequest) SetCreate(create int64) error {
    r.create = create
    r.Set("create", create)
    return nil
}

// Create Getter
func (r TaobaoQimenEventProduceRequest) GetCreate() int64 {
    return r.create
}
// Nick Setter
// 外部商家名称。必须同时填写platform
func (r *TaobaoQimenEventProduceRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoQimenEventProduceRequest) GetNick() string {
    return r.nick
}
