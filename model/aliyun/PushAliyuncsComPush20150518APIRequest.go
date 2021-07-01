package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* PushAliyuncsComPush20150518APIRequest
云推送指令API API请求
push.aliyuncs.com.push.20150518

阿里云推送新增API，允许一条推送指令同时发布到多个终端上。 */
type PushAliyuncsComPush20150518APIRequest struct {
	model.Params
	// 用户账号列表,以换行区分,仅sendType为3时有效
	_account string
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
	// 应用标识
	_appId int64
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 推送内容
	_body string
	// 推送接收设备，多个以逗号分隔
	_deviceId string
	// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下: 0:IOS设备; 1:Andriod设备 3:全部. 默认为3.
	_deviceType int64
	// iOS应用图标右上角角标
	_iOSBadge string
	// 自定义的kv结构,开发者扩展用 针对iOS设备
	_iOSExtParameters string
	// iOS通知声音
	_iOSMusic string
	// 推送时间,若空表示立即推送,推送时间不能早于当前时间
	_pushTime string
	// 当APP不在线时候，是否通过通知提醒
	_remind bool
	// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
	_sendType int64
	// 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到
	_storeOffline bool
	// 通知的摘要
	_summery string
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 推送的标题内容.
	_title string
}

// New
