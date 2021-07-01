package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* PushAliyuncsComPushNotification2015_03_18APIRequest
推送通知 API请求
push.aliyuncs.com.pushNotification.2015-03-18

pushNotification */
type PushAliyuncsComPushNotification2015_03_18APIRequest struct {
	model.Params
	// 用户账号列表,以换行区分,仅sendType为3时有效
	_account string
	// 自定义的kv结构,开发者扩展用
	_androidExtraMap string
	// 通知声音
	_androidMusic string
	// 通知类型 1:震动 2:响铃
	_androidNotifyType int64
	// 打开app指定位置
	_androidOpenActivity string
	// 点击通知后动作
	_androidOpenType int64
	// 打开应用,网页
	_androidOpenUrl string
	// 防打扰时长,取值范围为1~23
	_antiHarassDuration int64
	// 防打扰开始时间点,取值范围为0~23
	_antiHarassStartTime int64
	// 应用标识
	_appId int64
	// 批次编号,用于活动效果统计
	_batchNumber string
	// 设备编号列表,以换行区分,仅sendType为4时有效
	_deviceId string
	// 设备类型,取值范围为:0~3云推送支持多种设备,各 种设备类型编号如下:IOS设备:deviceType&amp;1=1; Andriod设备:deviceType&amp;2=2;如果存在此字段,则向 指定的设备类型推送消息。默认为全部(3);
	_deviceType int64
	// 自定义的kv结构,开发者扩展用
	_iosExtraMap string
	// 角标
	_iosFooter int64
	// 默认音乐
	_iosMusic string
	// 推送时间,若空表示立即推送,推送时间不能早于当前时间
	_pushTime string
	// 推送类型,取值范围:1~4; 1:所有人,无需指定tag、 deviceType等2:一群人,必须指定tag3:指定用户,根据 用户账号列表文件发送消息4:指定设备,根据设备编码列 表文件发送消息默认值为1
	_sendType int64
	// 摘要
	_summary string
	// 标签名称,仅支持1个标签,仅sendType为2时有效
	_tag string
	// 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	_timeout int64
	// 标题
	_title string
}

// New
