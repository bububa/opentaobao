package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusItemNotify 汽车票城际巴士车次变更通知飞猪接口
// taobao.bus.item.notify
//
// 供应商线路信息变更（如价格，发车时间，新增班次）同步到飞猪销售端需要 10分钟-10个小时的时间(跟供应商线路数量,接口可支持的并发量有关)。
// 为了让供应商线路信息变更能够及时体现到飞猪销售端，供应商可以在修改完自身系统的线路信息后，主动调用该接口通知飞猪，飞猪会将该线路数据进行一次同步。
func TaobaoBusItemNotify(clt *core.SDKClient, req *bus.TaobaoBusItemNotifyAPIRequest, resp *bus.TaobaoBusItemNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
