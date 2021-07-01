package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudpushPushAPIRequest
百川用户使用云推送高级推送接口 API请求
taobao.cloudpush.push

百川用户使用云推送高级推送接口 */
type TaobaoCloudpushPushAPIRequest struct {
	model.Params
	// 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
	_targetValue string
	// Android对应的activity,仅仅当androidOpenType=2有效
	_androidActivity string
	// 自定义的kv结构,开发者扩展用 针对android
	_androidExtParameters string
	// android通知声音
	_androidMusic string
	// 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
	_androidOpenType string
	// Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
	_androidOpenUrl string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 推送内容
	_body string
	// 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下:    iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
	_deviceType int64
	// iOS应用图标右上角角标
	_iosBadge string
	// 自定义的kv结构,开发者扩展用 针对iOS设备
	_iosExtParameters string
	// iOS通知声音
	_iosMusic string
	// 当APP不在线时候，是否通过通知提醒.  针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true  & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
	_remind bool
	// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
	_storeOffline bool
	// 通知的摘要
	_summery string
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 推送的标题内容.
	_title string
	// 0:表示消息(默认为0),1:表示通知
	_type int64
}

// New
