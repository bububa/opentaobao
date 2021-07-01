package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

/* TaobaoSmartstoreDeviceIashelfFeedback
智能硬件云货架数据回流
taobao.smartstore.device.iashelf.feedback

智慧门店云货架设备回流规则：（互动云货架、VR云货架通用）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.每一个action都必须传入用户操作时间op_time（start/end_time后续废弃）</br>
4.action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK时，item_id必须传入,且必须是淘宝商品的数字id</br>
5.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
6.outer_user 用于标识不能获取淘宝账号的游客</br> */
func TaobaoSmartstoreDeviceIashelfFeedback(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceIashelfFeedbackAPIRequest, session string) (*smartstore.TaobaoSmartstoreDeviceIashelfFeedbackAPIResponse, error) {
	var resp smartstore.TaobaoSmartstoreDeviceIashelfFeedbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
