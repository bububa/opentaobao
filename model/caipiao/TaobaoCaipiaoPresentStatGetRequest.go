package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家按天统计的彩票赠送数据 API请求
taobao.caipiao.present.stat.get

查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据.
*/
type TaobaoCaipiaoPresentStatGetRequest struct {
    model.Params
    // 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701.
    days   int64
}

// 初始化TaobaoCaipiaoPresentStatGetRequest对象
func NewTaobaoCaipiaoPresentStatGetRequest() *TaobaoCaipiaoPresentStatGetRequest{
    return &TaobaoCaipiaoPresentStatGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoPresentStatGetRequest) GetApiMethodName() string {
    return "taobao.caipiao.present.stat.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoPresentStatGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Days Setter
// 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701.
func (r *TaobaoCaipiaoPresentStatGetRequest) SetDays(days int64) error {
    r.days = days
    r.Set("days", days)
    return nil
}

// Days Getter
func (r TaobaoCaipiaoPresentStatGetRequest) GetDays() int64 {
    return r.days
}
