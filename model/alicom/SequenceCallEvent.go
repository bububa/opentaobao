package alicom

import (
	"sync"
)

// SequenceCallEvent 结构体
type SequenceCallEvent struct {
	// 呼出时间
	PCallOutTime string `json:"p_call_out_time,omitempty" xml:"p_call_out_time,omitempty"`
	// 呼出确认时间
	PCallAckTime string `json:"p_call_ack_time,omitempty" xml:"p_call_ack_time,omitempty"`
	// 响铃时间
	PRingTime string `json:"p_ring_time,omitempty" xml:"p_ring_time,omitempty"`
	// 释放时间
	PReleaseTime string `json:"p_release_time,omitempty" xml:"p_release_time,omitempty"`
	// 通话开始时间
	PStartTime string `json:"p_start_time,omitempty" xml:"p_start_time,omitempty"`
	// 呼叫结果
	PCallResult string `json:"p_call_result,omitempty" xml:"p_call_result,omitempty"`
	// 1
	PollingNo int64 `json:"polling_no,omitempty" xml:"polling_no,omitempty"`
	// 0-平台释放 1-主叫释放 2-被叫释放；短信话单时，传0
	PReleaseDir int64 `json:"p_release_dir,omitempty" xml:"p_release_dir,omitempty"`
	// 呼叫释放原因,1,未分配的号码（空号） 3,无至目的地的路由 4，停机 6,不可接受的信道 16,正常清除 17,用户忙 18,无用户响应 19,已有用户提醒，但无应答 21,呼叫拒绝 22,号码改变 26,清除未选择的用户 27,终点故障 28,无效号码格式（不完全的号码） 29,设施被拒绝 30,对状态询问的响应 31,正常--未规定 34,无电路/信道可用 38,网络故障 41,临时故障 42,交换设备拥塞 43,接入信息被丢弃 44,请求的电路/信道不可用 47,资源不可用--未规定 49,服务质量不可用 50,未预订所请求的设施 55,IncomingcallsbarredwithintheCUG 57,承载能力未认可(未开通通话功能） 58,承载能力目前不可用 63,无适用的业务或任选项目-未规定 65,承载业务不能实现 68,ACMequaltoorgreaterthanACMmax 69,所请求的设施不能实现 70,仅能获得受限数字信息承载能力 79,业务不能实现-未规定) 81,无效处理识别码 87UsernotmemberofCUG 88,非兼容目的地址 91,无效过渡网选择 95,无效消息-未规定 96,必选消息单元差错 97,消息类型不存在或不能实现 98,消息与控制状态不兼容-消息类型不存在或不能实现 99,信息单元不存在或不能实现 100,无效信息单元内容 101,消息与呼叫状态不兼容 102,定时器超时恢复 111,协议差错-未规定 127,互通-未规定 9999（短信话单时，传此值）
	PReleaseCause int64 `json:"p_release_cause,omitempty" xml:"p_release_cause,omitempty"`
}

var poolSequenceCallEvent = sync.Pool{
	New: func() any {
		return new(SequenceCallEvent)
	},
}

// GetSequenceCallEvent() 从对象池中获取SequenceCallEvent
func GetSequenceCallEvent() *SequenceCallEvent {
	return poolSequenceCallEvent.Get().(*SequenceCallEvent)
}

// ReleaseSequenceCallEvent 释放SequenceCallEvent
func ReleaseSequenceCallEvent(v *SequenceCallEvent) {
	v.PCallOutTime = ""
	v.PCallAckTime = ""
	v.PRingTime = ""
	v.PReleaseTime = ""
	v.PStartTime = ""
	v.PCallResult = ""
	v.PollingNo = 0
	v.PReleaseDir = 0
	v.PReleaseCause = 0
	poolSequenceCallEvent.Put(v)
}
