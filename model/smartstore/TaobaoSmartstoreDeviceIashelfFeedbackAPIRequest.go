package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest
智能硬件云货架数据回流 API请求
taobao.smartstore.device.iashelf.feedback

智慧门店云货架设备回流规则：（互动云货架、VR云货架通用）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time（start/end_time后续废弃）</br>
4.action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK时，item_id必须传入,且必须是淘宝商品的数字id</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br> */
type TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest struct {
	model.Params
	// 硬件CODE
	_deviceCode string
	// 字段废弃
	_endTime string
	// 字段废弃，考虑兼容，等同于op_time，两个必须传一个
	_startTime string
	// ACTION枚举值： BODY_SENSOR（通过人体感应、人脸识别成功识别到人） ITEM_SENSOR（通过RFID、蓝牙等感应设备识别商品，必须设置ITEM_ID） ITEM_CLICK（商品点击时必须设置ITEM_ID）  RECEIVE_COUPONS（扫码领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买，必须设置ITEM_ID）  SHARE_CLICK（点击分享）
	_action string
	// 商品ID，item_id 在action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK必须传入；  必须使用淘宝商品id，否则失败。
	_itemId string
	// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
	_couponId string
	// 数据外部编码，保证数据唯一性
	_outerBizId string
	// 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
	_opTime string
	// 用户昵称（混淆）
	_userNick string
	// 硬件识别的用户标识
	_outerUser string
}

// New
