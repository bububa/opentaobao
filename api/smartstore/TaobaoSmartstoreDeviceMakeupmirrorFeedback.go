package smartstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/smartstore"
)

/* 
智能硬件试妆镜数据回流 
taobao.smartstore.device.makeupmirror.feedback

智慧门店试妆镜设备回流规则（适用于试妆镜等）</br>
1.回流的数据属于当前授权的用户，回流的设备device_code由当前应用添加</br>
2.对于快闪店的智能硬件不需要授权</br>
3.action为SKU_CLICK时，sku_id必须传入</br>
4.action为ITEM_CLICK、SAMPLE_CLICK、BUY_CLICK、ITEM_FAVOR时，item_id必须传入,且必须是淘宝商品的数字id </br>
5.skin_detection 和scalp_detection 涉及相关检测功能的硬件设备回传 </br>
6.每一个acion都必须传入用户操作时间op_time </br>
7.outer_biz_id 用于硬件设备大量数据回流场景，服务商本地日志统计系统对一条日志记录生成唯一标识。 平台后端会对传入的outer_biz_id 做去重处理</br>
*/
func TaobaoSmartstoreDeviceMakeupmirrorFeedback(clt *core.SDKClient, req *smartstore.TaobaoSmartstoreDeviceMakeupmirrorFeedbackAPIRequest, session string) (*smartstore.TaobaoSmartstoreDeviceMakeupmirrorFeedbackAPIResponse, error) {
    var resp smartstore.TaobaoSmartstoreDeviceMakeupmirrorFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
